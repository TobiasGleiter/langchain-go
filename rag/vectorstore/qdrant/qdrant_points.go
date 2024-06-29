package qdrant

import (
	"context"
	"fmt"
	"net/http"
)

func (qs *QdrantStore) upsertPoints(ctx context.Context, payload UpsertPointsRequest) (UpsertPointsResponse, error) {
	fmt.Println("Upsert points")

	url := qs.Url.String() + "/collections/" + qs.Collection + "/points"
	fmt.Println(url)
	var response UpsertPointsResponse
	err := qs.sendHttpRequestWithContext(ctx, http.MethodPut, url, payload, &response)
	if err != nil {
		fmt.Println(err)
		return UpsertPointsResponse{}, err
	}
	return response, nil
}

func (qs *QdrantStore) deletePoints(ctx context.Context, payload DeletePointsRequest) (DeletePointsResponse, error) {
	url := qs.Url.String() + "/collections/" + qs.Collection + "/points/vectors/delete"
	var response DeletePointsResponse
	err := qs.sendHttpRequestWithContext(ctx, http.MethodPost, url, payload, &response)
	if err != nil {
		return DeletePointsResponse{}, err
	}
	return response, nil
}

func (qs *QdrantStore) countPoints(ctx context.Context, payload CountPointsRequest) (CountPointsResponse, error) {
	url := qs.Url.String() + "/collections/" + qs.Collection + "/points/count"
	var response CountPointsResponse
	err := qs.sendHttpRequestWithContext(ctx, http.MethodPost, url, payload, &response)
	if err != nil {
		return CountPointsResponse{}, err
	}
	return response, nil
}

func (qs *QdrantStore) retrievePoints(ctx context.Context, payload RetrievePointsRequest) (RetrievePointsResponse, error) {
	url := qs.Url.String() + "/collections/" + qs.Collection + "/points"
	var response RetrievePointsResponse
	err := qs.sendHttpRequestWithContext(ctx, http.MethodPost, url, payload, &response)
	if err != nil {
		return RetrievePointsResponse{}, err
	}
	return response, nil
}
