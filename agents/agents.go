package agents

import (
	"context"
	"errors"
)

var (
	ErrNotFinished          = errors.New("agent did not finish")
	ErrUnableToParseOutput  = errors.New("unable to parse output")
	ErrAgentNoReturn        = errors.New("agent did not return any action or finish")
	ErrExecutorInputNotString = errors.New("executor input not a string")
)

type Executor struct {
	Agent                   Agent
	MaxIterations           int
}

func NewExecutor(agent Agent) *Executor {
	return &Executor{Agent: agent, MaxIterations: 10}
}

func (e *Executor) Call(ctx context.Context, input string) any {
	return nil//e.getReturn(&AgentFinish{ReturnValues: make(map[string]any)}, steps), ErrNotFinished
}

// func (e *Executor) doIteration(
// 	ctx context.Context,
// 	steps []AgentStep,
// 	nameToTool map[string]tools.Tool,
// 	inputs map[string]string,
// ) ([]AgentStep, map[string]any, error) {
// 	actions, finish, err := e.Agent.Plan(ctx, steps, inputs)
// 	if errors.Is(err, ErrUnableToParseOutput) {
// 		steps = append(steps, AgentStep{
// 			Observation: err.Error(),
// 		})
// 		return steps, nil, nil
// 	}
// 	if err != nil {
// 		return steps, nil, err
// 	}

// 	if len(actions) == 0 && finish == nil {
// 		return steps, nil, ErrAgentNoReturn
// 	}

// 	if finish != nil {
// 		return steps, e.getReturn(finish, steps), nil
// 	}

// 	for _, action := range actions {
// 		steps, err = e.doAction(ctx, steps, nameToTool, action)
// 		if err != nil {
// 			return steps, nil, err
// 		}
// 	}

// 	return steps, nil, nil
// }