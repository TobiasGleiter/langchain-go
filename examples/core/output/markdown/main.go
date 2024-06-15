package main

import (
	"fmt"
	"io/ioutil"

	"github.com/TobiasGleiter/langchain-go/core/output"
)

func main() {
	content, _ := ioutil.ReadFile("markdown.md")
	markdown := string(content)

	var parser output.OutputParser[map[string]string] = &output.MarkdownOutputParser{}
	parsedMarkdownOutput, _ := parser.Parse(markdown)

	section := "## Heading 1.1"
	if content, exists := parsedMarkdownOutput[section]; exists {
		fmt.Printf("Content of %s:\n%s\n", section, content)
	} else {
		fmt.Printf("Section %s not found.\n", section)
	}
}
