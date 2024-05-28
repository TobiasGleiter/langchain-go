package main

import (
	"fmt"

	"github.com/TobiasGleiter/langchain-go/model_io/output"
)

type Joke struct {
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func main() {
	var parser output.OutputParser[Joke] = &output.JsonOutputParser[Joke]{}
	joke, err := parser.Parse(`{"setup": "Why don't scientists trust atoms?", "punchline": "Because they make up everything!"}`)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(joke.Setup)
	fmt.Println(joke.Punchline)
}