package main

import (
	"fmt"

	"github.com/TobiasGleiter/langchain-go/rag/similarity"
)

func main() {
	query := "This is a test query"
	document := "This query is a simple test"

	similarity := similarity.JaccardSimilarity(query, document)
	fmt.Printf("Jaccard Similarity: %.2f\n", similarity)
}
