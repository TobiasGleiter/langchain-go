package qdrant

import (
	"context"
	"net/url"

	"github.com/TobiasGleiter/langchain-go/core/embedder"
	"github.com/TobiasGleiter/langchain-go/rag/vectorstore"
)

// 1. Qdrant Client
// 2. Create a collection
// 3. Add Vectors
// 4. Run a query
// 5. Add a filter

type QdrantStore struct {
	Embedder       embedder.Embedder
	collectionName string
	qdrantURL      url.URL
	apiKey         string
	contentKey     string
}

func New(embedder embedder.Embedder) *QdrantStore {
	return &QdrantStore{
		Embedder: embedder,
	}
}

func (qs *QdrantStore) AddDocuments(ctx context.Context, docs []vectorstore.Document) ([]string, error) {
	return []string{}, nil
}

func (qs *QdrantStore) SimilaritySearch(ctx context.Context, query string, numberOfDocsToRetrun int) ([]vectorstore.Document, error) {
	return []vectorstore.Document{}, nil
}
