package qdrant

import (
	"context"
	"fmt"
	"net/url"

	"github.com/TobiasGleiter/langchain-go/core/embedder"
	"github.com/TobiasGleiter/langchain-go/rag/vectorstore"
	"github.com/google/uuid"
)

type QdrantStore struct {
	Embedder   embedder.Embedder
	Collection string
	Url        url.URL
	ApiKey     string
	contentKey string
}

func NewQdrant(embedder embedder.Embedder, collection string, url url.URL, apiKey string) *QdrantStore {
	return &QdrantStore{
		Embedder:   embedder,
		Collection: collection,
		Url:        url,
		ApiKey:     apiKey,
	}
}

func (qs *QdrantStore) AddDocuments(ctx context.Context, docs []vectorstore.Document) error {
	// 1. Create an array of documents
	texts := make([]string, 0, len(docs))
	for _, doc := range docs {
		texts = append(texts, doc.Content)
	}

	// 2. Create go routine for each index of the array and run embedding concurrently
	vectors := make([][]float32, 0, len(texts))
	for _, text := range texts {
		vector, _ := qs.Embedder.EmbedQuery(ctx, text)
		vectors = append(vectors, vector.Embedding)
	}

	// 3. Add metadata (what is the metadata?)
	metadatas := make([]map[string]interface{}, 0, len(docs))
	for i := 0; i < len(docs); i++ {
		metadata := make(map[string]interface{}, len(docs[i].Metadata))
		for key, value := range docs[i].Metadata {
			metadata[key] = value
		}
		metadatas = append(metadatas, metadata)
	}

	// 4. Create IDs for upsert points
	ids := make([]ID, len(vectors))
	for i := range ids {
		ids[i] = uuid.NewString()
	}

	// 5. Upsert points into qdrant vectorstore
	upsertPoints := UpsertPointsRequest{}
	upsertPoints.Batch.IDs = ids
	upsertPoints.Batch.Payloads = metadatas
	upsertPoints.Batch.Vectors = vectors

	response, err := qs.upsertPoints(ctx, upsertPoints) // make it only internal after bug fix
	if err != nil {
		return err
	}
	fmt.Println("Upsert Response", response)

	return nil
}

func (qs *QdrantStore) SimilaritySearch(ctx context.Context, query string) ([]string, error) {
	// 1. Embed query
	_, err := qs.Embedder.EmbedQuery(ctx, query)
	if err != nil {
		fmt.Println("Similarity search embedding error", err)
	}
	fmt.Println("Query embedded")

	// 2. Search similar points in vectorstore

	return []string{}, nil
}
