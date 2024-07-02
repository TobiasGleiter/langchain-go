package evaluator

type List = map[string]bool

type Evaluator interface {
	Parse(output string, config List) error
}
