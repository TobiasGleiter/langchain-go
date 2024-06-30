package qdrant

import (
	"context"
	"net/http"
)

func (qs *QdrantStore) createCollection(ctx context.Context, payload CreateCollectionRequest) (CreateCollectionResponse, error) {
	url := qs.Url.String() + "/collections/" + qs.Collection
	var response CreateCollectionResponse
	err := qs.sendHttpRequestWithContext(ctx, http.MethodPut, url, payload, &response)
	if err != nil {
		return CreateCollectionResponse{}, err
	}
	return response, nil
}

func (qs *QdrantStore) deleteCollection(ctx context.Context) (DeleteCollectionResponse, error) {
	url := qs.Url.String() + "/collections/" + qs.Collection
	var response DeleteCollectionResponse
	err := qs.sendHttpRequestWithContext(ctx, http.MethodDelete, url, "", &response)
	if err != nil {
		return DeleteCollectionResponse{}, err
	}
	return response, nil
}

func (qs *QdrantStore) CheckCollectionHealth(ctx context.Context) {}
