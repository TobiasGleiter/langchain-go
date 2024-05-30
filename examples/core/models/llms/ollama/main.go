package main

import (
	"fmt"
	"context"

	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
)

func main() {
	chatPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
        {Role: "system", Content: "You are a helpful assistant that translates {{.inputLanguage}} to {{.outputLanguage}}."},
        {Role: "user", Content: "{{.text}}"},
    })

	data := map[string]interface{}{
        "inputLanguage":  "English",
        "outputLanguage": "French",
        "text":           "I love programming.",
    }

	formattedMessages, err := chatPrompt.FormatMessages(data)
    if err != nil {
        panic(err)
    }
	
	llama3_8b_model := ollama.OllamaModel{
		Model:  "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
	}

	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)
	ctx := context.TODO()
	generatedContent, _ := ollamaClient.GenerateContent(ctx, formattedMessages)
	fmt.Println(generatedContent.Result)
}