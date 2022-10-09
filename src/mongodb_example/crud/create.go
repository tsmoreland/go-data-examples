package crud

import (
	"context"
	"github.com/tsmoreland/go-data-examples/mongodb_example/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func Add(collection *mongo.Collection, ctx context.Context, person model.Person) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(ctx, person)
}
