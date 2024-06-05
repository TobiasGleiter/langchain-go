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
	fmt.Println("Messages:", a.Messages)
	output, err := a.Model.GenerateContent(ctx, a.Messages)
	if err != nil {
		return AgentResponse{}, err
	}

	parts := strings.Split(output.Result, "Action:")
	var thought, action, tool, toolInput string

	fmt.Println("Parts",parts)

	if len(parts) == 2 {
		thoughtPart := strings.Split(parts[0], "Thought:")
		if len(thoughtPart) == 2 {
			thought = strings.TrimSpace(thoughtPart[1])
		} else {
			thought = "1 I should try again..."
		}

		actionParts := strings.Split(parts[1], "Action Input:")
		if len(actionParts) == 2 {
			action = strings.TrimSpace(actionParts[0])
			toolInput = strings.TrimSpace(actionParts[1])

			// This does not work correctly:
			if strings.HasPrefix(action, "[") && strings.HasSuffix(action, "]") {
				tool = strings.Trim(action, "[]")
				fmt.Println("[]", tool)
			} else {
				tool = action
			}
		} else {
			action = "2 I should try again..."
			toolInput = "None required."
			fmt.Println("Action Input part not found")
		}
	} else {
		// Reflection, because something went wrong generating or it is the final answer?!
		// thoughtParts := strings.Split(output.Result, "Thought:")
		// // thought = strings.TrimSpace(thoughtParts[1])
		// // fmt.Println("ohh.. \nThought:", thought)
		thought = "Did I find the answer?"

		// Thought will be there
		// Action is not: example:
		// ohh...  (GMT-5)
		// Thought: The current datetime indicates that today is Tuesday, June 4, 2024.
		// Final Answer: Today is a Tuesday.
		//fmt.Println("ohh...", output.Result) // Probably at the end because there is only a thought and Final Answer generated
	}

	
	if strings.Contains(output.Result, "Final Answer:") {
		finalAnswerParts := strings.Split(output.Result, "Final Answer:")
		finalAnswer := strings.TrimSpace(finalAnswerParts[1])
		fmt.Println("Final Answer:", finalAnswer)
		return AgentResponse{Finish: true}, nil
	}

	fmt.Println("Thought:", thought)
	fmt.Println("Action:", action)
	fmt.Println("Action Input:", toolInput)

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
				Content: "Thought: I cant find this tool. I should try one of these: [CurrentDatetime]",
			})
			return
		}

		observation, err := tool.Call(ctx, action.ToolInput)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// fmt.Println("Observation:", observation)

		a.Messages = append(a.Messages, models.MessageContent{
			Role: "assistant",
			Content: fmt.Sprintf("Observation: %s", observation),
		})
	}

	a.Actions = remainingActions // This removes all actions?
}

func parseToolString(toolString string) (string, string, error) {
	// Split the string by ": " to separate the index
	parts := strings.SplitN(toolString, ": ", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid format: %s", toolString)
	}
	
	// Extract the tool name and input from the remaining part
	toolWithInput := parts[1]
	if strings.Contains(toolWithInput, "[") && strings.Contains(toolWithInput, "]") {
		parts = strings.SplitN(toolWithInput, "[", 2)
		toolName := parts[0]
		toolInput := strings.TrimSuffix(parts[1], "]")
		return toolName, toolInput, nil
	}
	return "", "", fmt.Errorf("invalid tool format: %s", toolWithInput)
}