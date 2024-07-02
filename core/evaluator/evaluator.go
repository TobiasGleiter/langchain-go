package evaluator

import (
	"errors"
	"strings"
)

type WhitelistEvaluator struct{}

// Use lowercase words in config.
func (e *WhitelistEvaluator) Parse(output string, config Config) error {
	words := strings.Fields(output)
	for _, word := range words {
		if !config[word] {
			return errors.New("Not all words in whitelist.")
		}
	}
	return nil
}
