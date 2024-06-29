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
	// Embedding size of ollama3:8b = 4096
	llama3_8b_model := ollama.OllamaModel{
		Model:   "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream:  false,
	}

	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)
	ollamaEmbedder := embedder.NewEmbedder(ollamaClient)

	qdrantUrl := &url.URL{Scheme: "http", Host: "localhost:6333"}
	qdrant := qdrant.NewQdrant(*ollamaEmbedder, "collection_name", *qdrantUrl) // make sure the collection exists...
	//qdrant.UseApiKey("<api-key>")

	// docs := []vectorstore.Document{
	// 	{
	// 		Content:  "Take a leisurely walk in the park and enjoy the fresh air.",
	// 		Metadata: map[string]interface{}{"content": "Take a leisurely walk in the park and enjoy the fresh air."},
	// 	},
	// }

	ctx := context.Background()
	// err := qdrant.AddDocuments(ctx, docs)
	// if err != nil {
	// 	fmt.Println("Something went wrong...", err)
	// }

	similarPoints, _ := qdrant.SimilaritySearch(ctx, "I don't like to hike", 1)
	fmt.Println("Similar points:", similarPoints)
}
