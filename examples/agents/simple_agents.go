package main

import (
	"context"

	"github.com/TobiasGleiter/langchain-go/agents"
	"github.com/TobiasGleiter/langchain-go/tools"
)

type SimpleAgent struct {
	tools []tools.Tool
}

func (a *SimpleAgent) Plan(ctx context.Context, steps []agents.AgentStep, inputs map[string]string) ([]agents.AgentAction, *agents.AgentFinish, error) {
	// Check if we have completed any previous steps
	if len(steps) > 0 {
		// Produce a finish signal with the result
		finish := &agents.AgentFinish{
			ReturnValues: map[string]any{"output": steps[len(steps)-1].Observation},
		}
		return nil, finish, nil
	}

	// Otherwise, create a new action
	action := agents.AgentAction{Tool: "ECHO", ToolInput: inputs["input"]}
	return []agents.AgentAction{action}, nil, nil
}

func (a *SimpleAgent) GetTools() []tools.Tool {
	return a.tools
}

func (a *SimpleAgent) GetInputKeys() []string {
	return []string{"input"}
}

func (a *SimpleAgent) GetOutputKeys() []string {
	return []string{"output"}
}
