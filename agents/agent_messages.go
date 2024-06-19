package agents

import (
	"fmt"

	"github.com/TobiasGleiter/langchain-go/core/models"
)

func (a *Agent) addThoughtMessage(input string) {
	a.Messages = append(a.Messages, models.MessageContent{
		Role:    "assistant",
		Content: fmt.Sprintf("Thought: %s\n", input),
	})
}

func (a *Agent) addActionMessage(input string) {
	a.Messages = append(a.Messages, models.MessageContent{
		Role:    "assistant",
		Content: fmt.Sprintf("Action: %s\n", input),
	})
}

func (a *Agent) addActionInputMessage(input string) {
	a.Messages = append(a.Messages, models.MessageContent{
		Role:    "assistant",
		Content: fmt.Sprintf("Action Input: %s\n", input),
	})
}

func (a *Agent) addObservationMessage(observation string) {
	a.Messages = append(a.Messages, models.MessageContent{
		Role:    "assistant",
		Content: fmt.Sprintf("Observation: %s\n", observation),
	})
}

func (a *Agent) addObservationError(err error) {
	message := models.MessageContent{
		Role:    "assistant",
		Content: fmt.Sprintf("Observation: %s", err),
	}
	a.Messages = append(a.Messages, message)
}

func (a *Agent) addMissingTool() {
	tools := getToolNames(a.Tools)
	message := models.MessageContent{
		Role:    "assistant",
		Content: fmt.Sprintf("Thought: I can't find this tool. I should try one of these: [%s]\n", tools),
	}
	a.Messages = append(a.Messages, message)
}
