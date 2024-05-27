package main

import (
	"fmt"

	"github.com/TobiasGleiter/langchain-go/model_io/output"
)

func main() {
	stringOutputParser := output.StringOutputParser{}
	parsedStringOutput, _ := stringOutputParser.Parse("hi, bye")
	fmt.Println(parsedStringOutput)

	spaceSeparatedListOutputParser := output.SpaceSeparatedListOutputParser{}
	parsedSpaceSeparatedListOutput, _ := spaceSeparatedListOutputParser.Parse("hi, bye")
	fmt.Println(parsedSpaceSeparatedListOutput)

	commaSeparatedListOutputParser := output.CommaSeparatedListOutputParser{}
	parsedCommaSeparatedListOutput, _ := commaSeparatedListOutputParser.Parse("hi, bye")
	fmt.Println(parsedCommaSeparatedListOutput)
}