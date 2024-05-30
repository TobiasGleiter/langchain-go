package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"	

	"github.com/TobiasGleiter/langchain-go/core/models"
)

type OpenAiClient struct {
	APIKey   string
	ModelName string
}

func NewOpenAiClient(modelName, apiKey string) *OpenAiClient {
	return &OpenAiClient{
		APIKey:   apiKey,
		ModelName: modelName,
	}
}

func (oc *OpenAiClient) GenerateContent(ctx context.Context, messages []models.MessageContent) (models.ContentResponse, error) {
    httpClient := &http.Client{
        Timeout: 240 * time.Second,
    }

    requestPayload := OpenAIRequest{
        Model:       oc.ModelName,
        Messages:    messages,
        Temperature: 0.7,
    }

    requestBody, err := json.Marshal(requestPayload)
    if err != nil {
        return models.ContentResponse{}, err
    }

    req, err := http.NewRequestWithContext(ctx, "POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
    if err != nil {
        return models.ContentResponse{}, err
    }
    req.Header.Set("Authorization", "Bearer "+oc.APIKey)
    req.Header.Set("Content-Type", "application/json")

    resp, err := httpClient.Do(req)
    if err != nil {
        return models.ContentResponse{}, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return models.ContentResponse{}, err
    }

    var openAIResp OpenAIResponse
    if err := json.NewDecoder(resp.Body).Decode(&openAIResp); err != nil {
        return models.ContentResponse{}, err
    }

    if len(openAIResp.Choices) == 0 {
        return models.ContentResponse{}, err
    }

    contentResponse := models.ContentResponse{
        Result: openAIResp.Choices[0].Message.Content,
    }

    return contentResponse, nil
}
