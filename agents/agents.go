package agents

import (
	"fmt"
	"context"
	"strings"
	"regexp"

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

	fmt.Println(output.Result)

	parts := strings.Split(output.Result, fmt.Sprintf("\nAction "))
	var thought, action string
	if len(parts) == 2 {
		thought = strings.TrimSpace(parts[0])
		action = strings.TrimSpace(parts[1])
	} else {
		fmt.Println("ohh...", output.Result)
	}

	if strings.Contains(action, "Finish") {
		return AgentResponse{Actions: []AgentAction{}, Finish: true}, nil
	}

	a.Messages = append(a.Messages, models.MessageContent{
		Role: "assistant",
		Content: fmt.Sprintf("Thought %s\n", thought),
	})

	actions := []AgentAction{}
	tool, toolInput, err := parseToolString(action)
	if err != nil {
		return AgentResponse{Actions: []AgentAction{}, Finish: false}, nil
	}
	a.Actions = append(actions, AgentAction{
		Tool: tool,
		ToolInput: toolInput,
	})
	
	return AgentResponse{Actions: []AgentAction{}, Finish: false}, nil
}

func (a *Agent) Act(ctx context.Context) {
	var remainingActions []AgentAction

	for _, action := range a.Actions {
		tool, exists := a.Tools[action.Tool]
		if !exists {
			fmt.Println("Error: Tool not found")
		}

		observation, err := tool.Call(ctx, action.ToolInput)
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println("Observation:", observation)

		a.Messages = append(a.Messages, models.MessageContent{
			Role: "assistant",
			Content: observation,
		})
	}

	a.Actions = remainingActions // This removes all actions?
}

func extractFinishContent(input string) (string, error) {
	re := regexp.MustCompile(`Finish\[(.*?)\]`)
	matches := re.FindStringSubmatch(input)
	if len(matches) < 2 {
		return "", fmt.Errorf("no content found")
	}
	return matches[1], nil
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