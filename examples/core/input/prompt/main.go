package main

import (
	"fmt"

	"github.com/TobiasGleiter/langchain-go/model_io/input"
)

func main() {
	companyNamePrompt, _ := input.NewPromptTemplate("What is a good name for a company that makes {{.product}}?")

	data := map[string]any{ "product": "coloful socks" }
	companyNameFormattedPrompt, _ := companyNamePrompt.Format(data)
	fmt.Println(companyNameFormattedPrompt)

	buildProductPrompt, _ := prompts.NewPromptTemplate("{{.name}} want's to build {{.company}}.")

	buildProduct := map[string]any{ "name": "Tobi", "company": "coloful socks"}
	twoVariablesFormattedPrompt, _ := buildProductPrompt.Format(buildProduct)
	fmt.Println(twoVariablesFormattedPrompt)
}