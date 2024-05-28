package ollama

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
	"bytes"
	"io"

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

func (oc *OllamaClient) GenerateContent(ctx context.Context, messages []models.MessageContent) (models.ContentResponse, error) {
	client := &http.Client{
        Timeout: 240 * time.Second,
    }

	request := OllamaChatRequest{
		Model: oc.Model.Model,
		Messages: messages,
		Options: oc.Model.Options,
		Stream: oc.Model.Stream,
		Format: oc.Model.Format,
		KeepAlive: oc.Model.KeepAlive,
	}

    requestBody, err := json.Marshal(request)
    if err != nil {
        return models.ContentResponse{}, errors.New("error marshaling request")
    }

    req, err := http.NewRequest("POST", oc.Model.Endpoint, bytes.NewReader(requestBody))
    if err != nil {
        return models.ContentResponse{}, errors.New("create request failed")
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return models.ContentResponse{}, errors.New("HTTP request failed")
    }
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var chatResponse ChatResponse
	var finalResponse ChatResponse
	for {
		if err := decoder.Decode(&chatResponse); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return models.ContentResponse{}, errors.New("error decoding response")
		}

		finalResponse.Message.Content += chatResponse.Message.Content
	}

	if chatResponse.Done {
		return models.ContentResponse{Result: finalResponse.Message.Content}, nil	
	}

	

    return models.ContentResponse{Result: finalResponse.Message.Content}, nil	
}
