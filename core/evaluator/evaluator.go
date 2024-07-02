package evaluator

import (
	"errors"
	"strings"
)

type WordsWhitelistEvaluator struct{}
type WordsBlacklistEvaluator struct{}

func (e *WordsWhitelistEvaluator) Parse(output string, whitelist List) error {
	words := strings.Fields(output)
	for _, word := range words {
		if !whitelist[word] {
			return errors.New("Not all words in whitelist.")
		}
	}
	return nil
}

func (e *WordsBlacklistEvaluator) Parse(output string, blacklist List) error {
	words := strings.Fields(output)
	for _, word := range words {
		if blacklist[word] {
			return errors.New("Some words are in the blacklist.")
		}
	}
	return nil
}
