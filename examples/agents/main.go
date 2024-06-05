package main

import (
	"fmt"
	"time"
	"context"
	"io/ioutil"

	"github.com/TobiasGleiter/langchain-go/agents"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
)

// Implemented tools at the bottom of this file
type CurrentDatetime struct{}
type SaveToFile struct {}
type CurrentTemperature struct {}

func main() {
	// Agent need a llm and tools
	// Currently the agent works with the ReAct Prompt Pattern
	wizardlm2_7b := ollama.OllamaModel{
		Model:  "wizardlm2:7b", // This is the best working model, currently.
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
		Stop:   []string{"\nObservation", "Observation"}, // Necessary due to the ReAct Prompt Pattern
	}
	llm := ollama.NewOllamaClient(wizardlm2_7b)

	tools := map[string]agents.Tool{
		"CurrentDatetime": CurrentDatetime{},
		"CurrentTemperature": CurrentTemperature{},
		"SaveToFile": SaveToFile{},
	}

	agent := agents.NewAgent(llm, tools)
	agent.Task("How is the temperature? And save the temperature to a file: temperature.txt")

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

// Implementation of the Tools
func (t CurrentDatetime) Name() string {
	return "CurrentDatetime"
}

func (t CurrentDatetime) Call(ctx context.Context, input string) (string, error) {
	now := time.Now()
	formattedTime := now.Format(time.ANSIC)
	return fmt.Sprintf("Current datetime: %s", formattedTime), nil
}

func (t CurrentTemperature) Name() string {
	return "CurrentTemperature"
}

func (t CurrentTemperature) Call(ctx context.Context, input string) (string, error) {
	return fmt.Sprintf("Current temperature: 30°C"), nil
}

func (t SaveToFile) Name() string {
	return "SaveToFile"
}

func (t SaveToFile) Call(ctx context.Context, input string) (string, error) {
	// Save file logic here...
	// The tool can also be an agent (e.g. for formatting the output)
	data := []byte(input)
    err := ioutil.WriteFile("example.txt", data, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return "Error saving the file.", nil 
	}

	return "File saved.", nil
}