package embedder

import (
	"context"

	"github.com/TobiasGleiter/langchain-go/core/models"
)

type Embedder struct {
	LLM models.Model
}

func NewEmbedder(llm models.Model) *Embedder {
	return &Embedder{
		LLM: llm,
	}
}

func (e *Embedder) EmbedQuery(ctx context.Context, query string) (models.EmbeddingResponse, error) {
	return e.LLM.GenerateEmbedding(ctx, query)
}
