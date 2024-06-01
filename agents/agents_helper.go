package agents

// import (
// 	"fmt"
// 	"context"
// 	"errors"

// 	"github.com/TobiasGleiter/langchain-go/tools"
// )

// func (e *Executor) doAction(
// 	ctx context.Context,
// 	steps []AgentStep,
// 	nameToTool map[string]tools.Tool,
// 	action AgentAction,
// ) ([]AgentStep, error) {
// 	tool, ok := nameToTool[action.Tool]
// 	if !ok {
// 		return append(steps, AgentStep{
// 			Action:      action,
// 			Observation: fmt.Sprintf("%s is not a valid tool, try another one", action.Tool),
// 		}), nil
// 	}

// 	observation, err := tool.Call(ctx, action.ToolInput)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return append(steps, AgentStep{
// 		Action:      action,
// 		Observation: observation,
// 	}), nil
// }

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

// func inputsToString(inputValues map[string]any) (map[string]string, error) {
// 	inputs := make(map[string]string, len(inputValues))
// 	for key, value := range inputValues {
// 	  valueStr, ok := value.(string)
// 	  if !ok {
// 		return nil, fmt.Errorf("%w: %s", ErrExecutorInputNotString, key)
// 	  }
// 	  inputs[key] = valueStr
// 	}
// 	return inputs, nil
// }

// func getNameToTool(tools []tools.Tool) map[string]tools.Tool {
// 	if len(tools) == 0 {
// 		return nil
// 	}

// 	nameToTool := make(map[string]tools.Tool, len(tools))
// 	for _, tool := range tools {
// 		nameToTool[tool.Name()] = tool
// 	}

// 	return nameToTool
// }

// func (e *Executor) getReturn(finish *AgentFinish, steps []AgentStep) map[string]any {
// 	// if e.ReturnIntermediateSteps {
// 	// 	finish.ReturnValues["intermediateSteps"] = steps
// 	// }
// 	return finish.ReturnValues
// }