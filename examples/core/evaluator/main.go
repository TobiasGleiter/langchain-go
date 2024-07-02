package main

import (
	"fmt"

	"github.com/TobiasGleiter/langchain-go/core/evaluator"
)

func main() {
	outputFromModel := "You should try hiking instead."

	whitelist := evaluator.List{
		"You":      true,
		"should":   true,
		"try":      true,
		"hiking":   true,
		"instead.": true,
	}

	var whitelistEvaluator evaluator.Evaluator = &evaluator.WordsWhitelistEvaluator{}
	err := whitelistEvaluator.Parse(outputFromModel, whitelist)
	if err != nil {
		fmt.Println(err)
	}

	blacklist := evaluator.List{
		"You": true,
	}

	var blacklistEvaluator evaluator.Evaluator = &evaluator.WordsBlacklistEvaluator{}
	err = blacklistEvaluator.Parse(outputFromModel, blacklist)
	if err != nil {
		fmt.Println(err)
	}

}
