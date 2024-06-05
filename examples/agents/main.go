package main

import (
	"fmt"
	"time"
	"context"

	"github.com/TobiasGleiter/langchain-go/agents"
	"github.com/TobiasGleiter/langchain-go/core/input"
	"github.com/TobiasGleiter/langchain-go/core/models"
	"github.com/TobiasGleiter/langchain-go/core/models/llms/ollama"
)

type CurrentDatetime struct{}

func (t CurrentDatetime) Name() string {
	return "CurrentDatetime"
}

func (t CurrentDatetime) Call(ctx context.Context, input string) (string, error) {
	now := time.Now()
	formattedTime := now.Format(time.ANSIC)
	//fmt.Println("Tool is in use")
	return fmt.Sprintf("Current datetime: %s", formattedTime), nil
}

type SaveToFile struct {}

func (t SaveToFile) Name() string {
	return "SaveToFile"
}

func (t SaveToFile) Call(ctx context.Context, input string) (string, error) {
	// tools can be agents too?
	//fmt.Println("I saved the input to file", input)

	return "File saved to datetime.txt", nil
}

func main() {
	reActPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
        {Role: "user", Content: `
		Answer the following questions as best you can. You have access to the following external tools:

		[{{.tools}}]

		Use the following format:
		Question: the input question you must answer
		Thought: you should always think about what to do
		Action: the action to take, should be one of [{{.tools}}]
		Action Input: the input to the action
		Observation: the result of the action
		... (this Thought:/Action:/Action Input:/Observation: can repeat N times)
		Thought: I now know the final answer
		Final Answer: the final answer to the original input question

		`},
    })

	data := map[string]interface{}{
        "tools":		"CurrentDatetime, SaveToFile",
    }

	formattedMessages, err := reActPrompt.FormatMessages(data)
    if err != nil {
        panic(err)
    }

	wizardlm2_7b := ollama.OllamaModel{
		Model:  "wizardlm2:7b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
		Stop:   []string{"\nObservation", "Observation"},
	}
	llm := ollama.NewOllamaClient(wizardlm2_7b)


	tools := map[string]agents.Tool{
		"CurrentDatetime": CurrentDatetime{},
		"SaveToFile": SaveToFile{},
	}

	agent := agents.NewAgent(llm, tools, formattedMessages)
	agent.Task("What time is it? Save the time to a file.")

	iterationLimit := 10
	for i := 1; i < iterationLimit; i++ {
		ctx := context.TODO()
		todos, _ := agent.Plan(ctx)
		if todos.Finish {
			fmt.Printf("Finished")
			break
		}
		agent.Act(ctx) // Executes the actions from the plan (e.g. tools)
		fmt.Println(agent.Messages)
	}
}