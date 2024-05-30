package main

import (
	"fmt"
	"time"
	"log"
	"context"

	"github.com/TobiasGleiter/langchain-go/extension/txt"
	"github.com/TobiasGleiter/langchain-go/core/pipe"
	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/output"
	"github.com/TobiasGleiter/langchain-go/core/models"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
)

func main() {
	startTime := time.Now()

	textFileHandler := extension.NewTextFileHandler("example.txt")
	chunksOfText, err := textFileHandler.ReadChunks(2048)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(fmt.Sprintf("Chunks: %v", len(chunksOfText)))

	var summaries []string

	for _, chunk := range chunksOfText {

        summary := summarize(chunk)
        if err != nil {
            log.Printf("Error summarizing chunk: %v\n", err)
            continue // Skip to next chunk on error
        }

        summaries = append(summaries, summary)
    }

	var fullSummary string
	for _, summary := range summaries {
		fullSummary += summary
	}
	fullSummary = finalSummarize(fullSummary)

	err = textFileHandler.WriteAll(fullSummary)
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(startTime)
    fmt.Printf("Total execution time: %v\n", elapsed)
}

func summarize(text string) string {
	chatPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
        {Role: "system", Content: "Please provide a summary of the following text."},
        {Role: "user", Content: "TEXT: {{.text}}"},
    })

	data := map[string]interface{}{
        "text": text,
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
	var parser output.OutputParser[string] = &output.StringOutputParser{}

	ctx := context.TODO()
	pipe := pipe.NewPipe(messages, ollamaClient, parser)
	result, _ := pipe.Invoke(ctx)
	return result
}

func finalSummarize(text string) string {
	chatPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
        {Role: "system", Content: " Write a concise summary of the following text delimited by triple backquotes. Return your response in bullet points which covers the key points of the text."},
        {Role: "user", Content: "TEXT: {{.text}}"},
    })

	data := map[string]interface{}{
        "text": text,
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
	var parser output.OutputParser[string] = &output.StringOutputParser{}

	ctx := context.TODO()
	pipe := pipe.NewPipe(messages, ollamaClient, parser)
	result, _ := pipe.Invoke(ctx)
	return result
}