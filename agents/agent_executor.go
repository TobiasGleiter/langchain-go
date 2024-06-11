package agents

import (
	"fmt"
	"context"
)

type Executor struct {
	Agent *Agent
	IterationLimit int
}

func NewExecutor(agent *Agent) *Executor {
	e := &Executor{Agent: agent}
	e.IterationLimit = 10
	return e
}

func (e *Executor) Run(ctx context.Context) {
	for i := 1; i < e.IterationLimit; i++ {
		todos, _ := e.Agent.Plan(ctx)
		if todos.Finish {
			break
		}
		e.Agent.Act(ctx) // Executes the actions from the plan (e.g. tools)
	}
}

func (e *Executor) PrintMessages() {
    for idx := range e.Agent.Messages {
        fmt.Printf(e.Agent.Messages[idx].Content)
    }
    fmt.Println()	
}