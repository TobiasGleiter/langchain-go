package qdrant

import "context"

// 1. Qdrant Client
// 2. Create a collection https://qdrant.tech/documentation/concepts/collections/
// 3. Add Vectors
// 4. Run a query
// 5. Add a filter

func (qs *QdrantStore) CreateCollection(ctx context.Context, collection string) {
	// HTTP Request
	// 	PUT /collections/{collection_name}
	// {
	//     "vectors": {
	//       "size": 300,
	//       "distance": "Cosine"
	//     }
	// }
}

func (qs *QdrantStore) DeleteCollection(ctx context.Context, collection string) {
	// DELETE http://localhost:6333/collections/{collection_name}
}

func (qs *QdrantStore) CollectionHealth(ctx context.Context) {
	// GET http://localhost:6333/collections/{collection_name}/exists
}

func (qs *QdrantStore) CollectionInfo(ctx context.Context) {
	// GET /collections/{collection_name}
}
