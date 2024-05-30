package main

import (
	"fmt"

	"github.com/TobiasGleiter/langchain-go/model_io/output"
)

func main() {
	var parser output.OutputParser[string] = &output.StringOutputParser{}
	parsedStringOutput, _ := parser.Parse("hi, bye")
	fmt.Println(parsedStringOutput)

	var spaceParser output.OutputParser[[]string] = &output.SpaceSeparatedListOutputParser{}
    spaceResult, _ := spaceParser.Parse("example output with spaces")
    fmt.Println(spaceResult)

    var commaParser output.OutputParser[[]string] = &output.CommaSeparatedListOutputParser{}
    commaResult, _ := commaParser.Parse("example,output,with,commas")
    fmt.Println(commaResult)
}