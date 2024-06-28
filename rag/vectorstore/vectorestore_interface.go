package vectorstore

import "context"

type Map = map[string]any

type Document struct {
	PageContent string
	Metadata    Map
	Score       float32
}

type VectorStore interface {
	AddDocuments(ctx context.Context, docs []Document) ([]string, error)
	SimilaritySearch(ctx context.Context, query string, numberOfDocsToRetrun int) ([]Document, error)
}