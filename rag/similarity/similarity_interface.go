package similarity

type Similarity interface {
	Calculate(query, text string) (float64, error)
}
