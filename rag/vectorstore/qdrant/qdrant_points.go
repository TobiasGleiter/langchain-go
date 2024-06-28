package qdrant

import "context"

func (qs *QdrantStore) upsertPoints(ctx context.Context, vectors [][]float32) error {
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
	return nil
}

func (qs *QdrantStore) deletePoints(ctx context.Context, ids []any) error {
	//	POST /collections/{collection_name}/points/vectors/delete
	//
	//	{
	//	    "points": [0, 3, 100],
	//	    "vectors": ["text", "image"]
	//	}
	return nil
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
