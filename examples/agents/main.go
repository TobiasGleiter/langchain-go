package main

import (
	"fmt"
	"time"
	"context"

	"github.com/TobiasGleiter/langchain-go/agents"
	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
)

type CurrentDatetime struct{}

func (t CurrentDatetime) Name() string {
	return "CurrentDatetime"
}

func (t CurrentDatetime) Call(ctx context.Context, input string) (string, error) {
	now := time.Now()
	formattedTime := now.Format(time.ANSIC)
	fmt.Println("Tool is in use")
	return fmt.Sprintf("Current datetime: %s", formattedTime), nil
}

func main() {
	chatPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
        {Role: "user", Content: `
		Answer the following questions as best you can. You have access to the following tools:

		{{.tools}}

		Use the following format:
		Question: the input question you must answer
		Thought: you should always think about what to do
		Action: the action to take, should be one of [{{.tools}}]
		Action Input: the input to the action
		Obersvation: the result of the action
		... (this Thought/Action/Action Input/Observation can repeat N times)
		Thought: I now know the final answer
		Final Answer: the final answer to the original input question

		Begin!

		Question: {{.input}}
		Thought:
		`},
    })

	data := map[string]interface{}{
        "tools":		"CurrentDatetime",
        "input":		"What weekday is it and what month?",
    }

	formattedMessages, err := chatPrompt.FormatMessages(data)
    if err != nil {
        panic(err)
    }

	llama3_8b_model := ollama.OllamaModel{
		Model:  "wizardlm2:7b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
		Stop:   []string{"\nObservation", "Observation"},
	}
	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)


	tools := map[string]agents.Tool{
		"CurrentDatetime": CurrentDatetime{},
	}

	agent := agents.NewAgent(ollamaClient, tools, formattedMessages)

	iterationLimit := 10
	for i := 1; i < iterationLimit; i++ {
		ctx := context.TODO()
		todos, _ := agent.Plan(ctx) // Returns actions and finish
		if todos.Finish {
			fmt.Printf("Finished")
			break
		}
		agent.Act(ctx) // Executes the actions from the plan (e.g. tools)
		
	}
}