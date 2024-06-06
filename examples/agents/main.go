package main

import (
	"fmt"
	"time"
	"context"
	"io/ioutil"
	"strings"
	"strconv"

	"github.com/TobiasGleiter/langchain-go/agents"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
)

// Implemented tools at the bottom of this file
type CurrentDatetime struct{}
type CurrentTemperatureInFahrenheit struct {}
type FormatFahrenheitToCelsius struct {}
type SaveToFile struct {}

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
		"CurrentTemperatureInFahrenheit": CurrentTemperatureInFahrenheit{},
		"FormatFahrenheitToCelsius": FormatFahrenheitToCelsius{},
		"SaveToFile": SaveToFile{},
	}

	agent := agents.NewAgent(llm, tools)
	agent.Task("How is the temperature in celsius?")

	iterationLimit := 10
	for i := 1; i < iterationLimit; i++ {
		ctx := context.TODO()
		todos, _ := agent.Plan(ctx)
		if todos.Finish {
			fmt.Println(agent.Messages[len(agent.Messages)-1].Content)
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

func (t CurrentTemperatureInFahrenheit) Name() string {
	return "CurrentTemperatureInFahrenheit"
}

func (t CurrentTemperatureInFahrenheit) Call(ctx context.Context, input string) (string, error) {
	return fmt.Sprintf("Current temperature: 77°F"), nil
}

func (t FormatFahrenheitToCelsius) Name() string {
	return "FormatFahrenheitToCelsius"
}

func (t FormatFahrenheitToCelsius) Call(ctx context.Context, input string) (string, error) {
	fahrenheit, err := parseFahrenheit(input)
	if err != nil {
		return "", err
	}

	// Convert Fahrenheit to Celsius
	celsius := fahrenheitToCelsius(fahrenheit)

	// Return the formatted string
	return fmt.Sprintf("Current temperature: %.2f°C", celsius), nil
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

func parseFahrenheit(input string) (float64, error) {
	// Example input: "77°F"
	input = strings.TrimSpace(input)
	if !strings.HasSuffix(input, "°F") {
		return 0, fmt.Errorf("invalid input: %s", input)
	}

	// Remove the "°F" suffix
	value := strings.TrimSuffix(input, "°F")

	// Convert the string to a float64
	fahrenheit, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid temperature value: %s", value)
	}

	return fahrenheit, nil
}

// fahrenheitToCelsius converts Fahrenheit to Celsius
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}