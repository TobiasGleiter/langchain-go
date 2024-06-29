package qdrant

import (
	"context"
	"net/http"
)

func (qs *QdrantStore) CreateCollection(ctx context.Context, payload CreateCollectionRequest) (CreateCollectionResponse, error) {
	url := qs.Url.String() + "/collections/" + qs.Collection
	var response CreateCollectionResponse
	err := qs.sendHttpRequestWithContext(ctx, http.MethodPut, url, payload, &response)
	if err != nil {
		return CreateCollectionResponse{}, err
	}
	return response, nil
}
