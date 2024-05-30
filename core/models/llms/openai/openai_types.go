package openai

import (
	"github.com/TobiasGleiter/langchain-go/core/models"	
)

type OpenAIResponse struct {
	ID      string          `json:"id"`
	Object  string          `json:"object"`
	Created int64           `json:"created"`
	Model   string          `json:"model"`
	Usage   Usage
	Choices []OpenAIResponseChoice
}

type OpenAIResponseChoice struct {
	Message      struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
	Index        int         `json:"index"`
}

type Usage struct {
	PromptTokens      int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}


type OpenAIRequest struct {
	Model   string                    `json:"model"`
	Messages []models.MessageContent `json:"messages"`
	Temperature float32 `json:"temperature"`
}