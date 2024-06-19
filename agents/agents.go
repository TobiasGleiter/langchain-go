package agents

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models"
)

type Agent struct {
	Model    models.Model
	Tools    map[string]Tool
	Messages []models.MessageContent
	Actions  []AgentAction
}

func NewAgent(model models.Model, tools map[string]Tool) *Agent {
	toolNames := getToolNames(tools)
	initialMessages := setupReActPromptInitialMessages(toolNames)

	return &Agent{
		Model:    model,
		Tools:    tools,
		Messages: initialMessages,
	}
}

func (a *Agent) Task(prompt string) {
	chatPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
		{Role: "user", Content: `
			Begin!
			
			Question: {{.input}}
			Thought:
		`},
	})

	data := map[string]interface{}{
		"input": prompt,
	}

	formattedMessages, err := chatPrompt.FormatMessages(data)
	if err != nil {
		panic(err)
	}

	a.Messages = append(a.Messages, formattedMessages...)
}

func (a *Agent) Plan(ctx context.Context) (AgentResponse, error) {
	generatedContent, err := a.Model.GenerateContent(ctx, a.Messages)
	if err != nil {
		return AgentResponse{}, err
	}

	parts := strings.Split(generatedContent.Result, "Action:")
	var thought, action, tool, toolInput string

	if len(parts) == 2 {
		thoughtPart := strings.Split(parts[0], "Thought:")
		if len(thoughtPart) == 2 {
			thought = strings.TrimSpace(thoughtPart[1])
		} else {
			thought = "I should try again..."
		}

		actionParts := strings.Split(parts[1], "Action Input:")
		if len(actionParts) == 2 {
			action = strings.TrimSpace(actionParts[0])
			toolInput = strings.TrimSpace(actionParts[1])

			if strings.HasPrefix(action, "[") && strings.HasSuffix(action, "]") {
				tool = strings.Trim(action, "[]")
			} else {
				tool = action
			}
		} else {
			// Sometimes the model only outputs Thought and Action without Action Input.
			actionParts = strings.Split(parts[1], "Action:")
			tool = actionParts[0]
			toolInput = "None required."
		}

		// Ensure the message format.
		a.addThoughtMessage(thought)
		a.addActionMessage(action)
		a.addActionInputMessage(action)

		a.Actions = append([]AgentAction{}, AgentAction{
			Tool:      tool,
			ToolInput: toolInput,
		})
	}
	// TODO: else, handle wrong format...

	return a.CheckIfFinalAnswer(generatedContent.Result)
}

func (a *Agent) Act(ctx context.Context) {
	for _, action := range a.Actions {
		if !a.handleAction(ctx, action) {
			return
		}
	}
	a.clearActions()
}

func (a *Agent) handleAction(ctx context.Context, action AgentAction) bool {
	tool, exists := a.Tools[action.Tool]
	if !exists {
		a.addMissingTool()
		return false
	}

	observation, err := tool.Call(ctx, action.ToolInput)
	if err != nil {
		a.addObservationError(err)
		return false
	}

	a.addObservationMessage(observation)
	return true
}

func (a *Agent) clearActions() {
	a.Actions = nil
}

func (a *Agent) GetFinalAnswer() (string, error) {
	if len(a.Messages) == 0 {
		return "", errors.New("No messages provided")
	}
	finalAnswer := a.Messages[len(a.Messages)-1].Content
	parts := strings.Split(finalAnswer, "Final Answer: ")
	if len(parts) < 2 {
		return "", errors.New("Invalid final answer")
	}
	return parts[1], nil
}

func (a *Agent) CheckIfFinalAnswer(input string) (AgentResponse, error) {
	if strings.Contains(input, "FINAL ANSWER:") {
		finalAnswerParts := strings.Split(input, "FINAL ANSWER:")
		finalAnswer := strings.TrimSpace(finalAnswerParts[1])

		a.Messages = append(a.Messages, models.MessageContent{
			Role:    "assistant",
			Content: fmt.Sprintf("\nFinal Answer: %s", finalAnswer),
		})

		return AgentResponse{Finish: true}, nil
	}

	return AgentResponse{Finish: false}, nil
}

func setupReActPromptInitialMessages(tools string) []models.MessageContent {
	reActPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
		{Role: "user", Content: `
		Answer the following questions as best you can. 
		Use only values from the tools. Do not estimate or predict values.	
		Select the tool that fits the question:

		[{{.tools}}]

		Use the following format:

		Question: the input question you must answer
		Thought: you should always think about what to do
		Action: the action to take, should be one of [{{.tools}}]
		Action Input: the input to the action
		Observation: the result of the action
		... (this Thought:/Action:/Action Input:/Observation: can repeat N times)
		Thought: I now know the final answer

		FINAL ANSWER: the final answer to the original input question
		`},
	})

	data := map[string]interface{}{
		"tools": tools,
	}

	formattedMessages, err := reActPrompt.FormatMessages(data)
	if err != nil {
		panic(err)
	}

	return formattedMessages
}
