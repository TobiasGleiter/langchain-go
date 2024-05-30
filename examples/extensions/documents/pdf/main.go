package main

import (
	"fmt"
	"log"

	//"github.com/TobiasGleiter/langchain-go/extension/documents"
	"rsc.io/pdf"
)

func main() {
	pdfReader, err := pdf.Open("sample.pdf")
	if err != nil {
		// Handle error
	}
	numPages := pdfReader.NumPage()
	fmt.Println(numPages)

	for i := 1; i <= numPages; i++ {
		page := pdfReader.Page(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(page.Content)
	}
}