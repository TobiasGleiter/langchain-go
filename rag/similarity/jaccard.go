package similarity

import (
	"strings"

	"github.com/TobiasGleiter/langchain-go/internal/set"
)

type JaccardSimilarity struct{}

func (js JaccardSimilarity) Calculate(query, document string) (float64, error) {
	queryWords := strings.Fields(strings.ToLower(query))
	documentWords := strings.Fields(strings.ToLower(document))

	querySet := set.New[string]()
	documentSet := set.New[string]()

	for _, word := range queryWords {
		querySet.Add(word)
	}

	for _, word := range documentWords {
		documentSet.Add(word)
	}

	intersection := querySet.Intersection(documentSet)
	union := querySet.Union(documentSet)

	return float64(intersection.Size()) / float64(union.Size()), nil
}
