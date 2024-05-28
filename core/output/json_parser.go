package output

import (
	"encoding/json"
)

type JsonOutputParser[T any] struct{}

func (p *JsonOutputParser[T]) Parse(output string) (T, error) {
	var result T
	err := json.Unmarshal([]byte(output), &result)
	return result, err
}

func (p *JsonOutputParser[T]) ParseWithPrompt(output string, prompt PromptValue) (T, error) {
	// Implement logic if necessary, otherwise just parse the output
	return p.Parse(output)
}

func (p *JsonOutputParser[T]) GetFormatInstructions() string {
	return "Return the output as JSON."
}