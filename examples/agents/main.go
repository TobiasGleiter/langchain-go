package main

import (
	"fmt"
	"time"
	"context"

	"github.com/TobiasGleiter/langchain-go/agents"
	//"github.com/TobiasGleiter/langchain-go/core/input"
	//"github.com/TobiasGleiter/langchain-go/core/models"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
)

type CurrentDatetime struct{}

func (t CurrentDatetime) Name() string {
	return "CurrentDatetime"
}

func (t CurrentDatetime) Call(ctx context.Context, input string) (string, error) {
	now := time.Now()
	formattedTime := now.Format(time.ANSIC)
	//fmt.Println("Tool is in use")
	return fmt.Sprintf("Current datetime: %s", formattedTime), nil
}

type SaveToFile struct {}

func (t SaveToFile) Name() string {
	return "SaveToFile"
}

func (t SaveToFile) Call(ctx context.Context, input string) (string, error) {
	// tools can be agents too?
	//fmt.Println("I saved the input to file", input)

	return "File saved to datetime.txt", nil
}

func main() {

	wizardlm2_7b := ollama.OllamaModel{
		Model:  "wizardlm2:7b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
		Stop:   []string{"\nObservation", "Observation"},
	}
	llm := ollama.NewOllamaClient(wizardlm2_7b)


	tools := map[string]agents.Tool{
		"CurrentDatetime": CurrentDatetime{},
		"SaveToFile": SaveToFile{},
	}

	agent := agents.NewAgent(llm, tools)
	agent.Task("What time is it? Save the time to a file.")

	iterationLimit := 10
	for i := 1; i < iterationLimit; i++ {
		ctx := context.TODO()
		todos, _ := agent.Plan(ctx)
		if todos.Finish {
			fmt.Printf("Finished")
			break
		}
		agent.Act(ctx) // Executes the actions from the plan (e.g. tools)
		fmt.Println(agent.Messages)
	}
}