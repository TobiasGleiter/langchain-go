package ollama

import (
	"github.com/TobiasGleiter/langchain-go/core/models"	
)

type ChatResponse struct {
	Model				string		`json:"model"`
	CreatedAt			string		`json:"created_at"`
	Message 			models.MessageContent
	Done				bool		`json:"done"`
	TotalDuration      int64         `json:"total_duration"`
	LoadDuration       int64         `json:"load_duration"`
	PromptEvalCount    int           `json:"prompt_eval_count"`
	PromptEvalDuration int64         `json:"prompt_eval_duration"`
	EvalCount          int           `json:"eval_count"`
	EvalDuration       int64         `json:"eval_duration"`
}

type OllamaChatRequest struct {
    Model    string `json:"model"`
	Messages []models.MessageContent `json:"messages"`
    Options  ModelOptions
    Stream bool `json:"stream"`
	Format   string   `json:"format,omitempty"`
	KeepAlive int64 `json:"keepalive,omitempty"`	
}