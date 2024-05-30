package main

import (
	"fmt"
	
	"github.com/TobiasGleiter/langchain-go/extension/txt"
)

func main() {
	textFileHandler := extension.NewTextFileHandler("example.txt")

	content, err := textFileHandler.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = textFileHandler.WriteAll(content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Content saved to file successfully.")
}