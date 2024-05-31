package agents

import (
	"context"
	//"github.com/TobiasGleiter/langchain-go/tools"
)

type Agent interface {
	Plan(ctx context.Context, steps []AgentStep, inputs map[string]string) ([]AgentAction, *AgentFinish, error)
	GetTools() []Tool
	GetInputKeys() []string
	GetOutputKeys() []string
}
