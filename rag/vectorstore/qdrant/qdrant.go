package qdrant

import (
	"context"
	"net/url"

	"github.com/TobiasGleiter/langchain-go/core/embedder"
	"github.com/TobiasGleiter/langchain-go/rag/vectorstore"
)

type QdrantStore struct {
	Embedder   embedder.Embedder
	Collection string
	Url        url.URL
	apiKey     string
	contentKey string
}

func NewQdrant(embedder embedder.Embedder, collection string, url url.URL) *QdrantStore {
	return &QdrantStore{
		Embedder:   embedder,
		Collection: collection,
		Url:        url,
	}
}

func (qs *QdrantStore) AddDocuments(ctx context.Context, docs []vectorstore.Document) ([]string, error) {
	// 1. Create an array of documents

	// 2. Create go routine for each index of the array and run embedding concurrently

	// 3. Add metadata (what is the metadata?)

	// 4. Upsert points into qdrant vectorstore

	return []string{}, nil
}

func (qs *QdrantStore) SimilaritySearch(ctx context.Context, query string) ([]vectorstore.Document, error) {
	// 1. Embed query

	// 2. Search similar points in vectorstore

	return []vectorstore.Document{}, nil
}
