package qdrant

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (qs *QdrantStore) upsertPoints(ctx context.Context, payload UpsertPointsRequest) (UpsertPointsResponse, error) {
	upsertPoints, err := json.Marshal(payload)
	if err != nil {
		return UpsertPointsResponse{}, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	upsertPointsUrl := qs.Url.String() + "/collections/" + qs.Collection + "/points"
	upserPointsRequest, err := http.NewRequestWithContext(ctx, http.MethodPut, upsertPointsUrl, bytes.NewBuffer(upsertPoints))
	if err != nil {
		return UpsertPointsResponse{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	upserPointsRequest.Header.Set("Authorization", "Bearer "+qs.apiKey)
	upserPointsRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	upsertPointsResponse, err := client.Do(upserPointsRequest)
	if err != nil {
		return UpsertPointsResponse{}, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer upsertPointsResponse.Body.Close()

	if upsertPointsResponse.StatusCode != http.StatusOK {
		return UpsertPointsResponse{}, fmt.Errorf("unexpected status code: %d", upsertPointsResponse.StatusCode)
	}

	var upsertResponse UpsertPointsResponse
	err = json.NewDecoder(upsertPointsResponse.Body).Decode(&upsertResponse)
	if err != nil {
		return UpsertPointsResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return upsertResponse, nil
}

func (qs *QdrantStore) deletePoints(ctx context.Context, payload []DeletePointsRequest) (DeletePointsResponse, error) {
	deletePoints, err := json.Marshal(payload)
	if err != nil {
		return DeletePointsResponse{}, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	deletPointsUrl := qs.Url.String() + "/collections/" + qs.Collection + "/points/vectors/delete"
	deletePointsRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, deletPointsUrl, bytes.NewBuffer(deletePoints))
	if err != nil {
		return DeletePointsResponse{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	deletePointsRequest.Header.Set("Authorization", "Bearer "+qs.apiKey)
	deletePointsRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	deletePointsResponse, err := client.Do(deletePointsRequest)
	if err != nil {
		return DeletePointsResponse{}, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer deletePointsResponse.Body.Close()

	if deletePointsResponse.StatusCode != http.StatusOK {
		return DeletePointsResponse{}, fmt.Errorf("unexpected status code: %d", deletePointsResponse.StatusCode)
	}

	var deleteResponse DeletePointsResponse
	err = json.NewDecoder(deletePointsResponse.Body).Decode(&deleteResponse)
	if err != nil {
		return DeletePointsResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return deleteResponse, nil
}

func (qs *QdrantStore) countPoints(ctx context.Context, payload CountPointsRequest) (CountPointsResponse, error) {
	countPoints, err := json.Marshal(payload)
	if err != nil {
		return CountPointsResponse{}, fmt.Errorf("failed to marshal request payload: %w", err)
	}

	countPointsUrl := qs.Url.String() + "/collections/" + qs.Collection + "/points/count"
	countPointsRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, countPointsUrl, bytes.NewBuffer(countPoints))
	if err != nil {
		return CountPointsResponse{}, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	countPointsRequest.Header.Set("Authorization", "Bearer "+qs.apiKey)
	countPointsRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	countPointsResponse, err := client.Do(countPointsRequest)
	if err != nil {
		return CountPointsResponse{}, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer countPointsResponse.Body.Close()

	if countPointsResponse.StatusCode != http.StatusOK {
		return CountPointsResponse{}, fmt.Errorf("unexpected status code: %d", countPointsResponse.StatusCode)
	}

	var countResponse CountPointsResponse
	err = json.NewDecoder(countPointsResponse.Body).Decode(&countResponse)
	if err != nil {
		return CountPointsResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return countResponse, nil
}
