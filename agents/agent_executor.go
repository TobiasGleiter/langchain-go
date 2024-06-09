package agents

import (
	"fmt"
	"context"
)

type Executor struct {
	Agent *Agent
}

func NewExecutor(agent *Agent) *Executor {
	return &Executor{Agent: agent}
}

func (e *Executor) Run(ctx context.Context) {
	iterationLimit := 10
	lastShownIndex := 0

	for i := 1; i < iterationLimit; i++ {
		todos, _ := e.Agent.Plan(ctx)
		if todos.Finish {
			fmt.Println(e.Agent.Messages[len(e.Agent.Messages)-1].Content)
			fmt.Printf("Finished")
			break
		}
		e.Agent.Act(ctx) // Executes the actions from the plan (e.g. tools)

		fmt.Printf("Iteration %d: New messages:\n", i)
		for idx := lastShownIndex + 1; idx < len(e.Agent.Messages); idx++ {
			fmt.Printf(e.Agent.Messages[idx].Content)
		}
		fmt.Println()

		// Update the last shown index
		lastShownIndex = len(e.Agent.Messages) - 1
	}
}