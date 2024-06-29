package qdrant

type ID string

type UpsertPointsResponse struct {
	Time   float32            `json:"time"`
	Status string             `json:"status"`
	Result UpsertPointsResult `json:"result"`
}

type DeletePointsResponse = UpsertPointsResponse

type UpsertPointsResult struct {
	Status      string `json:"status"`
	OperationId int    `json:"operation_id"`
}

type UpsertPointsRequest struct {
	Batch struct {
		IDs      []ID                `json:"ids"`
		Payloads []map[string]string `json:"payloads"`
		Vectors  [][]float32         `json:"vectors"`
	} `json:"batch"`
}

type DeletePointsRequest struct {
	Points []ID
}

type CountPointsRequest struct{}

type CountPointsResponse struct {
	Time   float32 `json:"time"`
	Status string  `json:"status"`
	Result struct {
		Count int `json:"count"`
	} `json:"result"`
}
