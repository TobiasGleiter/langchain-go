package pipe

import (
	"context"

	"github.com/TobiasGleiter/langchain-go/core/output"
	"github.com/TobiasGleiter/langchain-go/core/models"
)

type Pipe[T any] struct {
	Messages     []models.MessageContent
	Model        models.Model
	OutputParser output.OutputParser[T]
}

func NewPipe[T any](messages []models.MessageContent, model models.Model, outputParser output.OutputParser[T]) *Pipe[T] {
	return &Pipe[T]{
		Messages:     messages,
		Model:        model,
		OutputParser: outputParser,
	}
}

func (p *Pipe[T]) Invoke(ctx context.Context) (T, error) {
	output, err := p.Model.GenerateContent(ctx, p.Messages)
	if err != nil {
		var zero T
		return zero, err
	}
	parsedOutput, err := p.OutputParser.Parse(output.Result)
	if err != nil {
		var zero T
		return zero, err
	}
	return parsedOutput, nil
}


