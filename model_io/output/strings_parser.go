// package output

// import (
// 	"strings"
// )

// type StringOutputParser struct {}
// type SpaceSeparatedListOutputParser struct {}
// type CommaSeparatedListOutputParser struct {}

// func (p *StringOutputParser) Parse(output string) (string, error) {
//     return output, nil
// }

// func (p *CommaSeparatedListOutputParser) Parse(output string) ([]string, error) {
//     return strings.Split(output, ","), nil
// }

// func (p *SpaceSeparatedListOutputParser) Parse(output string) ([]string, error) {
//     return strings.Split(output, " "), nil
// }

package output

import (
	"strings"
)

// Implemented String Parsers
type StringOutputParser struct{}
type SpaceSeparatedListOutputParser struct{}
type CommaSeparatedListOutputParser struct{}

func (p *StringOutputParser) Parse(output string) (string, error) {
	return output, nil
}

func (p *StringOutputParser) ParseWithPrompt(output string, prompt PromptValue) (string, error) {
	return output, nil
}

func (p *StringOutputParser) GetFormatInstructions() string {
	return "Return the output as a plain string."
}

func (p *SpaceSeparatedListOutputParser) Parse(output string) ([]string, error) {
	return strings.Split(output, " "), nil
}

func (p *SpaceSeparatedListOutputParser) ParseWithPrompt(output string, prompt PromptValue) ([]string, error) {
	return strings.Split(output, " "), nil
}

func (p *SpaceSeparatedListOutputParser) GetFormatInstructions() string {
	return "Return the output as a space-separated list."
}

func (p *CommaSeparatedListOutputParser) Parse(output string) ([]string, error) {
	return strings.Split(output, ","), nil
}

func (p *CommaSeparatedListOutputParser) ParseWithPrompt(output string, prompt PromptValue) ([]string, error) {
	return strings.Split(output, ","), nil
}

func (p *CommaSeparatedListOutputParser) GetFormatInstructions() string {
	return "Return the output as a comma-separated list."
}
