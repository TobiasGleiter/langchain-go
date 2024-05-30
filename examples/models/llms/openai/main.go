package main

import (
	"fmt"
	"log"
	"context"


	"github.com/TobiasGleiter/langchain-go/internal/env"
	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/openai"
)

func main() {
	env, err := env.LoadEnvFromPath("../../../../.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

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
	
	apiKey, ok := env["OPENAI_API_KEY"]
	if !ok {
		log.Fatal("OpenAI API key not found in .env file")
	}

	gpt_model := openai.Model{
		Model:  "gpt-3.5-turbo-0613",
		APIKey: apiKey,
	}

	openAiClient := openai.NewOpenAiClient(gpt_model)
	ctx := context.TODO()
	generatedContent, err := openAiClient.GenerateContent(ctx, formattedMessages)
	if err != nil {
		fmt.Println("Error generating content", err)
	}
	fmt.Println(generatedContent.Result)
}