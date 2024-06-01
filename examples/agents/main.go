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

func main() {
	agent := &SimpleAgent{}
	executor := agents.NewExecutor(agent)

	userInput := "Question: What is the elevation range for the area that the eastern sector of the Colorado orogeny extends into?" + " \n"
	ctx := context.Background()
	result := executor.Call(ctx, userInput)
	fmt.Println("Result:", result)
}