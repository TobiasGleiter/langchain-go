package main

import (
	"fmt"
	"log"

	"github.com/TobiasGleiter/langchain-go/extension/documents"
)

func main() {
	jsonReader := extension.NewJsonFileHandler("example.json")
	text, err := jsonReader.ReadAllToString()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(text)
}