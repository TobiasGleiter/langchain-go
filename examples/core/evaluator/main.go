package main

import (
	"log"

	"github.com/TobiasGleiter/langchain-go/core/evaluator"
)

func main() {
	outputFromModel := "You should try hiking instead."

	// Use only lower case words in config
	config := evaluator.Config{
		"You":      true,
		"should":   true,
		"try":      true,
		"hiking":   true,
		"instead.": true,
	}

	var whitelistEvaluator evaluator.Evaluator = &evaluator.WhitelistEvaluator{}
	err := whitelistEvaluator.Parse(outputFromModel, config)
	if err != nil {
		log.Fatal(err)
	}

}
