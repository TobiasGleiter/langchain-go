package main

import (
	"fmt"

	"github.com/TobiasGleiter/langchain-go/model_io/input/prompts"
)

func main() {
	companyNamePrompt, _ := prompts.NewPromptTemplate("What is a good name for a company that makes {{.product}}?")

	data := map[string]any{ "product": "coloful socks" }
	companyNameFormattedPrompt, _ := companyNamePrompt.Format(data)
	fmt.Println(companyNameFormattedPrompt)

	twoVariablesPrompt, _ := prompts.NewPromptTemplate("{{.name}} want's to build {{.company}}.")

	buildProduct := map[string]any{ "name": "Tobi", "company": "coloful socks"}
	twoVariablesFormattedPrompt, _ := twoVariablesPrompt.Format(buildProduct)
	fmt.Println(twoVariablesFormattedPrompt)
}