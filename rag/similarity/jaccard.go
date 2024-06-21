package similarity

import "strings"

func JaccardSimilarity(query string, document string) float64 {
	queryWords := strings.Fields(strings.ToLower(query))
	documentWords := strings.Fields(strings.ToLower(document))

	querySet := make(map[string]bool)
	documentSet := make(map[string]bool)

	for _, word := range queryWords {
		querySet[word] = true
	}

	for _, word := range documentWords {
		documentSet[word] = true
	}

	intersection := 0
	union := len(querySet)

	for word := range documentSet {
		if querySet[word] {
			intersection++
		} else {
			union++
		}
	}

	if union == 0 {
		return 0.0
	}
	return float64(intersection) / float64(union)
}
