package main

import (
	"fmt"

	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models"
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

	for _, message := range formattedMessages {
		fmt.Printf("[%s] %s\n", message.Role, message.Content)
	}
}
