package agents

import (
	"context"
	"errors"
	"fmt"
)

// Define an interface for tools.
type Tool interface {
	Name() string
	Call(ctx context.Context, input string) (string, error)
}

// Agent interface for planning and retrieving tools.


// Structures for agent actions and steps.
type AgentAction struct {
	Tool      string
	ToolInput string
}

type AgentStep struct {
	Action      AgentAction
	Observation string
}

type AgentFinish struct {
	ReturnValues map[string]any
}

var (
	ErrNotFinished          = errors.New("agent did not finish")
	ErrUnableToParseOutput  = errors.New("unable to parse output")
	ErrAgentNoReturn        = errors.New("agent did not return any action or finish")
	ErrExecutorInputNotString = errors.New("executor input not a string")
)

// Executor is the chain responsible for running agents.
type Executor struct {
	Agent                   Agent
	MaxIterations           int
	ReturnIntermediateSteps bool
}

// NewExecutor creates a new agent executor with an agent.
func NewExecutor(agent Agent) *Executor {
	return &Executor{Agent: agent, MaxIterations: 10, ReturnIntermediateSteps: false}
}

// Call executes the agent with the provided input values.
func (e *Executor) Call(ctx context.Context, inputValues map[string]any) (map[string]any, error) {
	inputs, err := inputsToString(inputValues)
	if err != nil {
		return nil, err
	}
	nameToTool := getNameToTool(e.Agent.GetTools())

	steps := make([]AgentStep, 0)
	for i := 0; i < e.MaxIterations; i++ {
		var finish map[string]any
		steps, finish, err = e.doIteration(ctx, steps, nameToTool, inputs)
		if finish != nil || err != nil {
			return finish, err
		}
	}

	return e.getReturn(&AgentFinish{ReturnValues: make(map[string]any)}, steps), ErrNotFinished
}

func (e *Executor) doIteration(
	ctx context.Context,
	steps []AgentStep,
	nameToTool map[string]Tool,
	inputs map[string]string,
) ([]AgentStep, map[string]any, error) {
	actions, finish, err := e.Agent.Plan(ctx, steps, inputs)
	if errors.Is(err, ErrUnableToParseOutput) {
		steps = append(steps, AgentStep{
			Observation: err.Error(),
		})
		return steps, nil, nil
	}
	if err != nil {
		return steps, nil, err
	}

	if len(actions) == 0 && finish == nil {
		return steps, nil, ErrAgentNoReturn
	}

	if finish != nil {
		return steps, e.getReturn(finish, steps), nil
	}

	for _, action := range actions {
		steps, err = e.doAction(ctx, steps, nameToTool, action)
		if err != nil {
			return steps, nil, err
		}
	}

	return steps, nil, nil
}

func (e *Executor) doAction(
	ctx context.Context,
	steps []AgentStep,
	nameToTool map[string]Tool,
	action AgentAction,
) ([]AgentStep, error) {
	tool, ok := nameToTool[action.Tool]
	if !ok {
		return append(steps, AgentStep{
			Action:      action,
			Observation: fmt.Sprintf("%s is not a valid tool, try another one", action.Tool),
		}), nil
	}

	observation, err := tool.Call(ctx, action.ToolInput)
	if err != nil {
		return nil, err
	}

	return append(steps, AgentStep{
		Action:      action,
		Observation: observation,
	}), nil
}

func (e *Executor) getReturn(finish *AgentFinish, steps []AgentStep) map[string]any {
	if e.ReturnIntermediateSteps {
		finish.ReturnValues["intermediateSteps"] = steps
	}
	return finish.ReturnValues
}

func inputsToString(inputValues map[string]any) (map[string]string, error) {
	inputs := make(map[string]string, len(inputValues))
	for key, value := range inputValues {
	  valueStr, ok := value.(string)
	  if !ok {
		return nil, fmt.Errorf("%w: %s", ErrExecutorInputNotString, key)
	  }
	  inputs[key] = valueStr
	}
	return inputs, nil
}

func getNameToTool(tools []Tool) map[string]Tool {
	if len(tools) == 0 {
		return nil
	}

	nameToTool := make(map[string]Tool, len(tools))
	for _, tool := range tools {
		nameToTool[tool.Name()] = tool
	}

	return nameToTool
}
