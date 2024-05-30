package main


import (
	"fmt"
	"context"

	"github.com/TobiasGleiter/langchain-go/core/pipe"
	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models"
	"github.com/TobiasGleiter/langchain-go/core/output"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
)

type Translation struct {
	Text string `json:"text"`
	InputLanguage string `json:"inputLanguage"`
	OutputLanguage string `json:"outputLanguage"`
}

func main() {
	chatPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
        {Role: "system", Content: "You are a helpful assistant that translates {{.InputLanguage}} to {{.OutputLanguage}}."},
        {Role: "user", Content: "{{.Text}}"},
    })

	data := Translation{
		Text: "I love programming.",
		InputLanguage: "English",
		OutputLanguage: "Spanish",
	}

	messages, err := chatPrompt.FormatMessages(data)
    if err != nil {
        panic(err)
    }
	
	llama3_8b_model := ollama.OllamaModel{
		Model:  "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
	}

	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)
	var parser output.OutputParser[Translation] = &output.JsonOutputParser[Translation]{}

	ctx := context.TODO()
	pipe := pipe.NewPipe(messages, ollamaClient, parser)
	result, _ := pipe.Invoke(ctx)
	fmt.Println(result.Text)
	fmt.Println(result.InputLanguage)
	fmt.Println(result.OutputLanguage)
	
}