package models

import (
	"context"
)

type MessageContent struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type ContentResponse struct {
	Result string
}

type Model interface {
	GenerateContent(ctx context.Context, messages []MessageContent) (ContentResponse, error)
}
