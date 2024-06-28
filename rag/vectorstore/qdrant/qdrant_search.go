package qdrant

type SearchPointsResponse struct {
	Result []SimilarItem `json:"result"`
	Status string        `json:"status"`
	Time   float32       `json:"time"`
}

type SimilarItem struct {
	Id    any     `json:"id"`
	Score float32 `json:"score"`
}

func (qs *QdrantStore) searchPoints(vector []float32) (SearchPointsResponse, error) {
	// 	POST /collections/{collection_name}/points/search
	// {
	//     "filter": {
	//         "must": [
	//             {
	//                 "key": "city",
	//                 "match": {
	//                     "value": "London"
	//                 }
	//             }
	//         ]
	//     },
	//     "params": {
	//         "hnsw_ef": 128,
	//         "exact": false
	//     },
	//     "vector": [0.2, 0.1, 0.9, 0.7],
	//     "limit": 3

	// Add the payload flag to retrieve the "human readable" payload
	// 	POST /collections/{collection_name}/points/search
	// {
	//     "vector": [0.2, 0.1, 0.9, 0.7],
	//     "with_vectors": true,
	//     "with_payload": true
	// }

	// Response:
	//
	//	{
	//		"result": [
	//		  { "id": 10, "score": 0.81 },
	//		  { "id": 14, "score": 0.75 },
	//		  { "id": 11, "score": 0.73 }
	//		],
	//		"status": "ok",
	//		"time": 0.001
	//	  }
	//	}
	return SearchPointsResponse{}, nil
}
