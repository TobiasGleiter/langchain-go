package main

import (
	"context"
	"fmt"
	"github.com/TobiasGleiter/langchain-go/agents"
)

type EchoTool struct{}

func (e *EchoTool) Name() string {
	return "ECHO"
}

func (e *EchoTool) Call(ctx context.Context, input string) (string, error) {
	return "Echo: " + input, nil
}

type SimpleAgent struct {
	tools []agents.Tool
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

func (a *SimpleAgent) GetTools() []agents.Tool {
	return a.tools
}

func (a *SimpleAgent) GetInputKeys() []string {
	return []string{"input"}
}

func (a *SimpleAgent) GetOutputKeys() []string {
	return []string{"output"}
}

func main() {
	agent := &SimpleAgent{
		tools: []agents.Tool{&EchoTool{}},
	}
	executor := agents.NewExecutor(agent)

	ctx := context.Background()
	inputValues := map[string]any{"input": "Hello, world!"}
	result, err := executor.Call(ctx, inputValues)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}