package agents

import (
	"fmt"
	"context"
	"strings"

	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models"
)

type Agent struct {
	Model models.Model
	Tools map[string]Tool
	Messages []models.MessageContent
	Actions []AgentAction
}

func NewAgent(model models.Model, tools map[string]Tool) *Agent {
	toolNames := getToolNames(tools)
	initialMessages := setupReActPromptInitialMessages(toolNames)

	return &Agent{
		Model: model,
		Tools: tools,
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
        "input":		prompt,
    }

	formattedMessages, err := chatPrompt.FormatMessages(data)
    if err != nil {
        panic(err)
    }

	a.Messages = append(a.Messages, formattedMessages...)
}

func (a *Agent) Plan(ctx context.Context) (AgentResponse,  error) {	
	output, err := a.Model.GenerateContent(ctx, a.Messages)
	if err != nil {
		return AgentResponse{}, err
	}

	parts := strings.Split(output.Result, "Action:")
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
			fmt.Println("ACtion:", actionParts)
			tool = actionParts[0]
			toolInput = "None required."
		}
	} else {
		// When this entered, the agent loop to a rather poor result
		thought = "I should try again..."
		action = "No action required?"
		toolInput = "No input required?"
	}

	
	if strings.Contains(output.Result, "Final Answer:") {
		finalAnswerParts := strings.Split(output.Result, "Final Answer:")
		finalAnswer := strings.TrimSpace(finalAnswerParts[1])

		a.Messages = append(a.Messages, models.MessageContent{
			Role: "assistant",
			Content: fmt.Sprintf("Final Answer: %s\n", finalAnswer),
		})

		return AgentResponse{Finish: true}, nil
	}

	// Ensure the message format.
	a.Messages = append(a.Messages, models.MessageContent{
		Role: "assistant",
		Content: fmt.Sprintf("Thought: %s\n", thought),
	})

	a.Messages = append(a.Messages, models.MessageContent{
		Role: "assistant",
		Content: fmt.Sprintf("Action: %s\n", action),
	})

	a.Messages = append(a.Messages, models.MessageContent{
		Role: "assistant",
		Content: fmt.Sprintf("Action Input: %s\n", toolInput),
	})


	actions := []AgentAction{}
	if len(action) > 0 {
		a.Actions = append(actions, AgentAction{
			Tool: tool,
			ToolInput: toolInput,
		})
	}


	return AgentResponse{Finish: false}, nil
}

func (a *Agent) Act(ctx context.Context) {
	var remainingActions []AgentAction

	for _, action := range a.Actions {
		tool, exists := a.Tools[action.Tool]
		if !exists {
			tools := getToolNames(a.Tools)
			a.Messages = append(a.Messages, models.MessageContent{
				Role: "assistant",
				Content: fmt.Sprintf("Thought: I cant find this tool. I should try one of these: [%s]", tools),
			})
			return
		}

		observation, err := tool.Call(ctx, action.ToolInput)
		if err != nil {
			a.Messages = append(a.Messages, models.MessageContent{
				Role: "assistant",
				Content: fmt.Sprintf("Observation: %s", err),
			})
			fmt.Println("Error:", err)
			return
		}

		a.Messages = append(a.Messages, models.MessageContent{
			Role: "assistant",
			Content: fmt.Sprintf("Observation: %s", observation),
		})
	}

	a.Actions = remainingActions // This removes all actions?
}

func getToolNames(tools map[string]Tool) string {
	names := make([]string, 0, len(tools))
	for _, tool := range tools {
		names = append(names, tool.Name())
	}
	return strings.Join(names, ", ")
}

//You have access to the following external tools:
func setupReActPromptInitialMessages(tools string) []models.MessageContent {
	reActPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
        {Role: "user", Content: `
		Answer the following questions as best you can. Do not estimate or predict values.
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

		Final Answer: the final answer to the original input question

		`},
    })

	data := map[string]interface{}{
        "tools":		tools,
    }

	formattedMessages, err := reActPrompt.FormatMessages(data)
    if err != nil {
        panic(err)
    }
	
	return formattedMessages
}
