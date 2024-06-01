package main

import (
	"fmt"
	"context"

	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
	"github.com/TobiasGleiter/langchain-go/agents"
)


func main() {
	llama3_8b_model := ollama.OllamaModel{
		Model:  "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
	}
	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)


}

	// agent := agents.NewAgent(ollamaClient)

	// ctx := context.TODO()
	// output := agent.Plan(ctx, "Hallo")
	// fmt.Println(output)