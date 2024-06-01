package agents

import (
	"context"

	"github.com/TobiasGleiter/langchain-go/core/models"
)

type Tool interface {
	Name() string
	Call(input string) (string, error)
}

type Agent struct {
	Model models.Model
	Messages []models.MessageContent
}

func NewAgent(model models.Model) *Agent {
	return &Agent{
		Model: model,
	}
}

func (agent *Agent) Plan(ctx context.Context, input string) string {
	message := models.MessageContent{
		Role: "user",
		Content: input,
	}
	agent.Messages = append(agent.Messages, message)

	output, err := agent.Model.GenerateContent(ctx, agent.Messages)
	if err != nil {
		return "error generating plan"
	}
	return output.Result
}