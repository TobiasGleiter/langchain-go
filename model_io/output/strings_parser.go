package output

import (
	"strings"
)

type StringOutputParser struct {}
type SpaceSeparatedListOutputParser struct {}
type CommaSeparatedListOutputParser struct {}

func (p *StringOutputParser) Parse(output string) (string, error) {
    return output, nil
}

func (p *CommaSeparatedListOutputParser) Parse(output string) ([]string, error) {
    return strings.Split(output, ","), nil
}

func (p *SpaceSeparatedListOutputParser) Parse(output string) ([]string, error) {
    return strings.Split(output, " "), nil
}
