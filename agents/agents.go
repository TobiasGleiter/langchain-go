package agents

import (
	"fmt"
	"context"
	"strings"

	//"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models"
)

type Agent struct {
	Model models.Model
	Tools map[string]Tool
	Messages []models.MessageContent
	Actions []AgentAction
}

func NewAgent(model models.Model, tools map[string]Tool, messages []models.MessageContent) *Agent {
	return &Agent{
		Model: model,
		Tools: tools,
		Messages: messages,
	}
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
			action = "I should try again..."
			toolInput = "None required."
		}
	} else {
		thought = "Did I find the answer?"
		action = ""
		toolInput = ""
	}

	
	if strings.Contains(output.Result, "Final Answer:") {
		finalAnswerParts := strings.Split(output.Result, "Final Answer:")
		finalAnswer := strings.TrimSpace(finalAnswerParts[1])
		fmt.Println("Final Answer:", finalAnswer)
		return AgentResponse{Finish: true}, nil
	}

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
	a.Actions = append(actions, AgentAction{
		Tool: tool,
		ToolInput: toolInput,
	})

	return AgentResponse{Finish: false}, nil
}

func (a *Agent) Act(ctx context.Context) {
	var remainingActions []AgentAction

	for _, action := range a.Actions {
		tool, exists := a.Tools[action.Tool]
		if !exists {
			a.Messages = append(a.Messages, models.MessageContent{
				Role: "assistant",
				Content: "Thought: I cant find this tool. I should try one of these: [CurrentDatetime, SaveToFile]",
			})
			return
		}

		observation, err := tool.Call(ctx, action.ToolInput)
		if err != nil {
			fmt.Println("Error:", err)
		}

		a.Messages = append(a.Messages, models.MessageContent{
			Role: "assistant",
			Content: fmt.Sprintf("Observation: %s", observation),
		})
	}

	a.Actions = remainingActions // This removes all actions?
}