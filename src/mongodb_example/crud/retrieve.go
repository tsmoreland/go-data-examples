package crud

import (
	"context"
	"github.com/tsmoreland/go-data-examples/mongodb_example/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOneByFilter(collection *mongo.Collection, ctx context.Context, filter interface{},
	opts ...*options.FindOneOptions) (*model.Person, error) {
	var person model.Person
	err := collection.FindOne(ctx, filter, opts...).Decode(&person)
	if err != nil {
		return nil, err
	} else {
		return &person, nil
	}
}
