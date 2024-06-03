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
	fmt.Println("Tool is in use")
	return fmt.Sprintf("Current datetime: %s", formattedTime), nil
}

func main() {
	chatPrompt, _ := input.NewChatPromptTemplate([]models.MessageContent{
        {Role: "user", Content: `
		Answer the following questions as best you can. You have access to the following tools:

		{{.tools}}

		Use the following format:
		Question: the input question you must answer
		Thought: you should always think about what to do
		Action: the action to take, should be one of [{{.tool_names}}]
		Action Input: the input to the action
		Obersvation: the result of the action
		... (this Thought/Action/Action Input/Observation can repeat N times)
		Thought: I now know the final answer
		Final Answer: the final answer to the original input question

		Begin!

		Question: {{.input}}
		Thought: 
		`},
    })

	data := map[string]interface{}{
        "tools":		"CurrentDatetime",
        "tool_names":	"CurrentDatetime",
        "input":		"What day is it?",
    }

	formattedMessages, err := chatPrompt.FormatMessages(data)
    if err != nil {
        panic(err)
    }

	// userInput := "Question: What day is it?" + " \n"

	// instruction := `Solve a question answering task with interleaving Thought, Action, Observation steps. 
	// 	Thought can reason about the current situation, and Action can be two types: 
	// 	(1) CurrentDatetime[entity], which returns the current time of the device.
	// 	(3) Finish[answer], which returns the full answer and finishes the task.
	// 	Only use these actions. Here are some examples.
	// 	`

	// example := `
	// Question: Musician and satirist Allie Goertz wrote a song about the \"The Simpsons\" character Milhouse, who Matt Groening named after who?\nThought 1: I only need to search Milhouse and find who it is named after.\nAction 1: Search[Milhouse]\nObservation 1: Milhouse Mussolini Van Houten is a recurring character in the Fox animated television series The Simpsons voiced by Pamela Hayden and created by Matt Groening. Milhouse is Bart Simpson's best friend in Mrs. Krabappel's fourth grade class at Springfield Elementary School. He is an insecure, gullible, and less popular child than Bart who is often led into trouble by Bart, who takes advantage of his friend's na\u00c3\u00afvet\u00c3\u00a9. Milhouse is a regular target for school bully Nelson Muntz and his friends Jimbo Jones, Dolph Starbeam and Kearney Zzyzwicz. Milhouse has a crush on Bart's sister, Lisa, a common plot element.\nThought 2: The paragraph does not tell who Milhouse is named after, maybe I can look up \"named after\".\nAction 2: Lookup[named after]\nObservation 2: (Result 1 / 1) Milhouse was designed by Matt Groening for a planned series on NBC, which was abandoned.[4] The design was then used for a Butterfinger commercial, and it was decided to use the character in the series.[5][6] Milhouse was named after U.S. president Richard Nixon, whose middle name was Milhous. The name was the most \"unfortunate name Matt Groening could think of for a kid\".[1] Years earlier, in a 1986 Life in Hell comic entitled \"What to Name the Baby\", Groening listed Milhouse as a name \"no longer recommended\".[7] Milhouse is a favorite among the staff as Al Jean noted \"most of the writers are more like Milhouse than Bart\".[1] His last name was given to him by Howard Gewirtz, a freelance writer who wrote the episode \"Homer Defined\"
	// `

	// var prompt = userInput
	// var messages []models.MessageContent
	// messages = append(messages, models.MessageContent{
	// 	Role: "system",
	// 	Content: fmt.Sprintf(`%s %s`, instruction, example),
	// })

	// messages = append(messages, models.MessageContent{
	// 	Role: "user",
	// 	Content: prompt+fmt.Sprintf("Thought 1:"),
	// })

	llama3_8b_model := ollama.OllamaModel{
		Model:  "wizardlm2:7b",
		Options: ollama.ModelOptions{NumCtx: 4096},
		Stream: false,
		Stop:   []string{"\nObservation", "Observation"},
	}
	ollamaClient := ollama.NewOllamaClient(llama3_8b_model)


	tools := map[string]agents.Tool{
		"CurrentDatetime": CurrentDatetime{},
	}

	
	agent := agents.NewAgent(ollamaClient, tools, formattedMessages)

	iterationLimit := 10
	for i := 1; i < iterationLimit; i++ {
		ctx := context.TODO()
		todos, _ := agent.Plan(ctx) // Returns actions and finish
		if todos.Finish {
			fmt.Printf("Finished")
			break
		}
		agent.Act(ctx) // Executes the actions from the plan (e.g. tools)
		
	}
}