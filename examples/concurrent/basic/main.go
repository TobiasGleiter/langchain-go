package main

import (
	"context"
	"fmt"
	"time"

	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
)

func worker(messages []models.MessageContent, ch chan string, name string) {
	llama3_8b_model := ollama.OllamaModel{
		Model:   "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream:  false,
	}

	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)
	ctx := context.TODO()
	output, _ := ollamaClient.GenerateContent(ctx, messages)

	fmt.Println(output)

	ch <- name
}

func main() {
	start := time.Now()

	chatPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
		{Role: "system", Content: "You are a helpful assistant that translates {{.inputLanguage}} to {{.outputLanguage}}."},
		{Role: "user", Content: "{{.text}}"},
	})

	data := map[string]interface{}{
		"inputLanguage":  "English",
		"outputLanguage": "Spanish",
		"text":           "I love coding!",
	}

	messages, err := chatPrompt.FormatMessages(data)
	if err != nil {
		panic(err)
	}

	c1 := make(chan string)
	c2 := make(chan string)

	go worker(messages, c1, "one")
	go worker(messages, c2, "two")

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Total execution time: %s\n", elapsed)
}
