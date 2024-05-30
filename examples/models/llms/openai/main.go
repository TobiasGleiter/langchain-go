package main

import (
	"fmt"
	"context"

	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/openai"
)

func main() {
	chatPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
        {Role: "user", Content: "{{.text}}"},
    })

	data := map[string]interface{}{
        "text":           "This is a test!",
    }

	formattedMessages, err := chatPrompt.FormatMessages(data)
    if err != nil {
        panic(err)
    }

	fmt.Println(formattedMessages)
	
	apiKey := ""
	modelName := "gpt-3.5-turbo"
	openAiClient := openai.NewOpenAiClient(modelName, apiKey)
	ctx := context.TODO()
	generatedContent, err := openAiClient.GenerateContent(ctx, formattedMessages)
	if err != nil {
		fmt.Println("Error generating content", err)
	}
	fmt.Println(generatedContent.Result)
}