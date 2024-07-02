package evaluator

type Config = map[string]bool

type Evaluator interface {
	Parse(text string) error
}
