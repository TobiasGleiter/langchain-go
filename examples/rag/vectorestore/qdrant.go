package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/TobiasGleiter/langchain-go/core/embedder"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
	"github.com/TobiasGleiter/langchain-go/rag/vectorstore"
	"github.com/TobiasGleiter/langchain-go/rag/vectorstore/qdrant"
)

func main() {
	/*
			curl -X PUT http://localhost:6333/collections/collection_name \
		     -H "api-key: <apiKey>" \
		     -H "Content-Type: application/json" \
		     -d '{
		            "vectors": {
		                "size": 4096,
		                "distance": "Cosine"
		            }
		        }'
	*/

	// Embedding size of ollama3:8b = 4096
	llama3_8b_model := ollama.OllamaModel{
		Model:   "llama3:8b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream:  false,
	}

	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)
	ollamaEmbedder := embedder.NewEmbedder(ollamaClient)

	qdrantUrl, _ := url.Parse("http://localhost:6333")
	qdrant := qdrant.NewQdrant(*ollamaEmbedder, "collection_name", *qdrantUrl)

	docs := []vectorstore.Document{
		{
			Content: "Take a leisurely walk in the park and enjoy the fresh air.",
			Metadata: map[string]any{
				"activities": "walking",
			},
		},
	}

	ctx := context.Background()
	qdrant.AddDocuments(ctx, docs)
	res, _ := qdrant.SimilaritySearch(ctx, "I don't think...")
	fmt.Println(res)
}
