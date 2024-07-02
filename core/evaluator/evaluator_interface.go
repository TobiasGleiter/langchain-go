package evaluator

type Config = map[string]bool

type Evaluator interface {
	Parse(output string, config Config) error
}
