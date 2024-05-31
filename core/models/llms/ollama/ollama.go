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

	var endpoint = oc.Model.Endpoint
	if oc.Model.Endpoint == "" {
		endpoint = "http://localhost:11434/api/chat"
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

    req, err := http.NewRequest("POST", endpoint, bytes.NewReader(requestBody))
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

func (oc *OllamaClient) GenerateEmbedding(ctx context.Context, prompt string) (models.EmbeddingResponse, error) {
	client := &http.Client{
        Timeout: 240 * time.Second,
    }

	var endpoint = oc.Model.Endpoint
	if oc.Model.Endpoint == "" {
		endpoint = "http://localhost:11434/api/embeddings"
	}

	request := OllamaPromptRequest{
		Model: oc.Model.Model,
		Prompt: prompt,
		Options: oc.Model.Options,
		Format: oc.Model.Format,
		KeepAlive: oc.Model.KeepAlive,
	}

    requestBody, err := json.Marshal(request)
    if err != nil {
        return models.EmbeddingResponse{}, errors.New("error marshaling request")
    }

    req, err := http.NewRequest("POST", endpoint, bytes.NewReader(requestBody))
    if err != nil {
        return models.EmbeddingResponse{}, errors.New("create request failed")
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return models.EmbeddingResponse{}, errors.New("HTTP request failed")
    }
	defer resp.Body.Close()

	var embeddingResponse models.EmbeddingResponse
	err = json.NewDecoder(resp.Body).Decode(&embeddingResponse)
	if err != nil {
		return models.EmbeddingResponse{}, errors.New("error decoding response")
	}

	

    return models.EmbeddingResponse{Embedding: embeddingResponse.Embedding}, nil	
}