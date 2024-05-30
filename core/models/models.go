package models

import (
	"context"
)

type MessageContent struct {
	Role string
	Content string
}

type ContentResponse struct {
	Result string
}

type Model interface {
	GenerateContent(ctx context.Context, messages []MessageContent) (ContentResponse, error)
}
