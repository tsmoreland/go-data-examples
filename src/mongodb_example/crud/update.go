package crud

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func Update(collection *mongo.Collection, ctx context.Context, filter interface{},
	update interface{}) (int64, int64, error) {

	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, 0, err
	} else {
		return updateResult.MatchedCount, updateResult.ModifiedCount, nil
	}

}
