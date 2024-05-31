package main

import (
	"fmt"
	"context"

	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
)

func main() {
	companyNamePrompt, _ := input.NewPromptTemplate("What is a good name for a company that makes {{.product}}?")

	data := map[string]any{ "product": "coloful socks" }
	companyNameFormattedPrompt, _ := companyNamePrompt.Format(data)
	
	llama3_8b_model := ollama.OllamaModel{
		Model:  "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
	}

	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)
	ctx := context.TODO()
	generatedContent, _ := ollamaClient.GenerateEmbedding(ctx, companyNameFormattedPrompt)
	fmt.Println(len(generatedContent.Embedding))
}