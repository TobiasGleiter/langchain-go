package main

import (
	"fmt"

	"github.com/TobiasGleiter/langchain-go/rag/similarity"
)

func main() {
	corpus := []string{
		"Take a leisurely walk in the park and enjoy the fresh air.",
		"Visit a local museum and discover something new.",
		"Attend a live music concert and feel the rhythm.",
		"Go for a hike and admire the natural scenery.",
		"Have a picnic with friends and share some laughs.",
		"Explore a new cuisine by dining at an ethnic restaurant.",
		"Take a yoga class and stretch your body and mind.",
		"Join a local sports league and enjoy some friendly competition.",
		"Attend a workshop or lecture on a topic you're interested in.",
		"Visit an amusement park and ride the roller coasters.",
	}

	query := "I like to hike"
	// Just for showcase, should be refactored to be used with different similarity searches.
	similarDocument := selectOneSimilarDocument(query, corpus) // Complexity of O(n)
	fmt.Println(similarDocument)

}

func selectOneSimilarDocument(query string, corpus []string) string {
	maxSimilarity := -1.0
	bestMatchIndex := -1
	for i, doc := range corpus {
		similarity := similarity.JaccardSimilarity(query, doc)
		if similarity > maxSimilarity {
			maxSimilarity = similarity
			bestMatchIndex = i
		}
	}

	return corpus[bestMatchIndex]
}
