package similarity

type Similarity[T any] interface {
	Calculate(query, document T) (float64, error)
}
