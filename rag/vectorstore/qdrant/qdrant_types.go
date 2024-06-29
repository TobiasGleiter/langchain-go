package qdrant

type ID string
type Map = map[string]string

type UpsertPointsRequest struct {
	Batch struct {
		IDs      []ID        `json:"ids"`
		Payloads []Map       `json:"payloads"`
		Vectors  [][]float32 `json:"vectors"`
	} `json:"batch"`
}

type UpsertPointsResponse struct {
	Time   float32            `json:"time"`
	Status string             `json:"status"`
	Result UpsertPointsResult `json:"result"`
}

type UpsertPointsResult struct {
	Status      string `json:"status"`
	OperationId int    `json:"operation_id"`
}

type DeletePointsRequest struct {
	Points []ID
}

type DeletePointsResponse = UpsertPointsResponse

type CountPointsRequest struct{}

type CountPointsResponse struct {
	Time   float32 `json:"time"`
	Status string  `json:"status"`
	Result struct {
		Count int `json:"count"`
	} `json:"result"`
}

type SearchPointsRequest struct {
}

type SearchPointsResponse struct {
	Result []SimilarItem `json:"result"`
	Status string        `json:"status"`
	Time   float32       `json:"time"`
}

type SimilarItem struct {
	Id      any     `json:"id"`
	Score   float32 `json:"score"`
	Payload Map     `json:"payload"`
}
