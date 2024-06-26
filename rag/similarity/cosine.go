package similarity

import (
	"math"
)

// Formula: (x.y) / (||x||*||y||)
// x.y (dot product between x and y) => Sum of i to n for xi * yi
// ||x|| is sqrt(x*x)

type Cosine struct{}

func (cs Cosine) Calculate(query, vector []float64) (float64, error) {
	dotProduct := 0.0
	queryMagnitude := 0.0
	vectorMagnitude := 0.0

	for i := range vector {
		dotProduct += vector[i] * query[i]
		queryMagnitude += query[i] * query[i]
		vectorMagnitude += vector[i] * vector[i]
	}

	return dotProduct / (math.Sqrt(queryMagnitude) * math.Sqrt(vectorMagnitude)), nil
}
