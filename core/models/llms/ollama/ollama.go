package ollama

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"bytes"

	"github.com/TobiasGleiter/langchain-go/core/models"
)

type ModelOptions struct {
	NumCtx int `json:"num_ctx"`
	Temperature float64 `json:"temperature"`
}

// Specific message and response structures for Ollama
type OllamaMessageContent struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OllamaContentResponse struct {
	Result  string `json:"result"`
	Details string `json:"details"`
}

// OllamaModel struct
type OllamaModel struct {
	Model     string       `json:"model"`
	APIKey    string       `json:"api_key"`
	Endpoint  string       `json:"endpoint"`
	Options   ModelOptions `json:"options"`
	Stream    bool         `json:"stream"`
	Format    string       `json:"format,omitempty"`
	KeepAlive int64        `json:"keepalive,omitempty"`
	Stop      []string     `json:"stop"`
}

type OllamaClient struct {
	Model OllamaModel
}

func NewOllamaClient(model OllamaModel) *OllamaClient {
	return &OllamaClient{Model: model}
}

func (oc *OllamaClient) GenerateContent(ctx context.Context, messages []models.MessageContent) (*models.ContentResponse, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Convert general messages to specific Ollama messages
	var ollamaMessages []OllamaMessageContent
	for _, msg := range messages {
		ollamaMessages = append(ollamaMessages, OllamaMessageContent{Role: msg.Role, Content: msg.Content})
	}

	payload, err := json.Marshal(ollamaMessages)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", oc.Model.Endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+oc.Model.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to generate content: " + resp.Status)
	}

	var ollamaResponse OllamaContentResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResponse); err != nil {
		return nil, err
	}

	return &models.ContentResponse{Result: ollamaResponse.Result}, nil
}
