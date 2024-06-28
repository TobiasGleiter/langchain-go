package qdrant

import "context"

type UpsertPointsResponse struct {
	Time   float32            `json:"time"`
	Status string             `json:"status"`
	Result UpsertPointsResult `json:"result"`
}

type DeletePointsResult = UpsertPointsResponse

type UpsertPointsResult struct {
	Status      string `json:"status"`
	OperationId int    `json:"operation_id"`
}

func (qs *QdrantStore) upsertPoints(ctx context.Context, vectors [][]float32) (UpsertPointsResponse, error) {
	// PUT /collections/{collection_name}/points
	// {
	// 	"batch": {
	// 		"ids": [1, 2, 3],
	// 		"payloads": [
	// 			{"color": "red"},
	// 			{"color": "green"},
	// 			{"color": "blue"}
	// 		],
	// 		"vectors": [
	// 			[0.9, 0.1, 0.1],
	// 			[0.1, 0.9, 0.1],
	// 			[0.1, 0.1, 0.9]
	// 		]
	// 	}
	// }
	return UpsertPointsResponse{}, nil
}

func (qs *QdrantStore) deletePoints(ctx context.Context, ids []any) (DeletePointsResult, error) {
	//	POST /collections/{collection_name}/points/vectors/delete
	//
	//	{
	//	    "points": [0, 3, 100],
	//	    "vectors": ["text", "image"]
	//	}
	return DeletePointsResult{}, nil
}

func (qs *QdrantStore) countPoints() (int, error) {
	//	POST /collections/{collection_name}/points/count
	//
	//	{
	//	    "filter": {
	//	        "must": [
	//	            {
	//	                "key": "color",
	//	                "match": {
	//	                    "value": "red"
	//	                }
	//	            }
	//	        ]
	//	    },
	//	    "exact": true
	//	}
	return 0, nil
}
