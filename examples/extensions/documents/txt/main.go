package main

import (
	"fmt"
	
	"github.com/TobiasGleiter/langchain-go/extension/documents"
)

func main() {
	textFileHandler := extension.NewTextFileHandler("example.txt")

	content, err := textFileHandler.ReadAllToString()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = textFileHandler.WriteContentToFile(content, "example_save.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Content saved to file successfully.")
}