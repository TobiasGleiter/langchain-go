package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/TobiasGleiter/langchain-go/core/embedder"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
	"github.com/TobiasGleiter/langchain-go/rag/vectorstore/qdrant"
)

func main() {
	llama3_8b_model := ollama.OllamaModel{
		Model:   "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream:  false,
	}

	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)
	ollamaEmbedder := embedder.NewEmbedder(ollamaClient)

	qdrantUrl, _ := url.Parse("http://localhost:6333")
	qdrant := qdrant.NewQdrant(*ollamaEmbedder, "collection", *qdrantUrl)

	ctx := context.Background()
	res, _ := qdrant.SimilaritySearch(ctx, "I don't think...")
	fmt.Println(res)
}
