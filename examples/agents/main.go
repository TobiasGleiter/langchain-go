package main

import (
	"fmt"
	"context"
	"strings"
	"strconv"
	"errors"

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
	// Agent Executor iterates 10 timees that the agent can solve the task.
	// Currently the agent works with the ReAct Prompt Pattern
	wizardlm2_7b := ollama.OllamaModel{
		Model:  "wizardlm2:7b", // This is the best working model, currently.
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
		Stop:   []string{"\nObservation", "Observation"}, // Necessary due to the ReAct Prompt Pattern
	}
	llm := ollama.NewOllamaClient(wizardlm2_7b)

	tools := map[string]agents.Tool{
		"CurrentTemperatureInFahrenheit": CurrentTemperatureInFahrenheit{},
		"FormatFahrenheitToCelsius": FormatFahrenheitToCelsius{},
	}

	agent := agents.NewAgent(llm, tools)
	agent.Task("How is the temperature in celsius?")

	ctx := context.TODO()
	executor := agents.NewExecutor(*agent)
	executor.Run(ctx)

}

// Implementation of the tools!
// Tools follow the tools interface with Name and Call.
func (t CurrentTemperatureInFahrenheit) Name() string {
	return "CurrentTemperatureInFahrenheit"
}

func (t CurrentTemperatureInFahrenheit) Call(ctx context.Context, input string) (string, error) {
	return fmt.Sprintf("Current temperature: 112°F"), nil
}

func (t FormatFahrenheitToCelsius) Name() string {
	return "FormatFahrenheitToCelsius"
}

func (t FormatFahrenheitToCelsius) Call(ctx context.Context, input string) (string, error) {
	fahrenheit, err := parseFahrenheit(input)
	if err != nil {
		// This need to be given by the user as an extra description that the model can reflect itself.
		return "", errors.New("I should use this tool again using the temperature in this format: XX°F and pass it to the Action Input:")
	}

	celsius := fahrenheitToCelsius(fahrenheit)

	return fmt.Sprintf("Current temperature: %.2f°C", celsius), nil
}

func parseFahrenheit(input string) (float64, error) {
	input = strings.TrimSpace(input)

	var value string

	if strings.HasSuffix(input, "°F") {
		// Remove the "°F" suffix
		value = strings.TrimSuffix(input, "°F")
	} else if strings.HasSuffix(input, "degrees Fahrenheit") {
		// Remove the "degrees Fahrenheit" suffix
		value = strings.TrimSuffix(input, "degrees Fahrenheit")
		value = strings.TrimSpace(value)
	} else {
		return 0, fmt.Errorf("invalid input: %s", input)
	}

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