package main

import (
	"fmt"

	"github.com/TobiasGleiter/langchain-go/rag/similarity"
)

func main() {
	var cosineSimilarity = &similarity.Cosine{} // Complexity O(n)

	score, _ := cosineSimilarity.Calculate([]float64{1, 1, 1}, []float64{1, 1, 1})
	fmt.Println(score)

	score, _ = cosineSimilarity.Calculate([]float64{1, 0, 1}, []float64{0, 1, 0})
	fmt.Println(score)

	score, _ = cosineSimilarity.Calculate([]float64{1, 0, 1}, []float64{1, 1, 0})
	fmt.Println(score)
}
