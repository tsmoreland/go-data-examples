package crud

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(collection *mongo.Collection, ctx context.Context, filter interface{}) (int64, error) {
	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	} else {
		return deleteResult.DeletedCount, nil
	}

}
