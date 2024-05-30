package main

import (
	"fmt"
	"github.com/TobiasGleiter/langchain-go/extension/input"
)

func main() {
	content, err := extension.ReadTextFile("example")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = extension.SaveToTextFile("example_output", content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Content saved to file successfully.")
}