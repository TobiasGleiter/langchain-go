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
	ApiKey     string
	contentKey string
}

func NewQdrant(embedder embedder.Embedder, collection string, url url.URL) *QdrantStore {
	return &QdrantStore{
		Embedder:   embedder,
		Collection: collection,
		Url:        url,
		ApiKey:     "",
	}
}

func (qs *QdrantStore) UseApiKey(apiKey string) {
	qs.ApiKey = apiKey
}

func (qs *QdrantStore) AddDocuments(ctx context.Context, docs []vectorstore.Document) (string, error) {
	upsertPointIds := qs.createUpsertPointIds(docs)
	metadatas := qs.createMetadatas(docs)
	contents := qs.createDocumentContent(docs)
	embeddedContents := qs.embedDocumentContents(ctx, contents)

	upsertPoints := UpsertPointsRequest{}
	upsertPoints.Batch.IDs = upsertPointIds
	upsertPoints.Batch.Payloads = metadatas
	upsertPoints.Batch.Vectors = embeddedContents

	response, err := qs.upsertPoints(ctx, upsertPoints)
	if err != nil {
		return response.Status, err
	}

	return response.Status, nil
}

func (qs *QdrantStore) SimilaritySearch(ctx context.Context, query string, limit int) ([]string, error) {
	// 1. Embed query
	embeddedResponse, err := qs.Embedder.EmbedQuery(ctx, query)
	if err != nil {
		return nil, err
	}

	// 2. Search similar points in vectorstore
	payload := SearchPointsRequest{
		Limit:       limit,
		WithPayload: true,
		WithVector:  false,
		Vector:      embeddedResponse.Embedding,
	}

	similarPoints, err := qs.searchPoints(ctx, payload)
	if err != nil {
		return nil, err
	}

	// 3. Retrieve content field from payload
	var similarPointsContent []string
	for _, similarPoint := range similarPoints.Result {
		similarPointsContent = append(similarPointsContent, similarPoint.Payload["content"].(string))
	}

	return similarPointsContent, nil
}
