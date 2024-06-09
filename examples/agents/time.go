package main

import (
	"context"
	"time"

	"github.com/TobiasGleiter/langchain-go/agents"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
)

type GetTime struct {}

func main() {
	wizardlm2_7b := ollama.OllamaModel{
		Model:  "wizardlm2:7b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
		Stop:   []string{"\nObservation", "Observation"},
	}
	llm := ollama.NewOllamaClient(wizardlm2_7b)


	tools := map[string]agents.Tool{
		"GetTime": GetTime{},
	}

	timeAgent := agents.NewAgent(llm, tools)
	timeAgent.Task("What time is it?")

	ctx := context.TODO()
	executor := agents.NewExecutor(timeAgent)
	executor.Run(ctx)
	executor.PrintMessages()
}

func (t GetTime) Name() string { return "GetTime" }

func (t GetTime) Call(ctx context.Context, input string) (string, error) {
	currentTime := time.Now()
	return currentTime.String(), nil
}