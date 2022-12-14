package infrastructure

import (
	"context"
	"fmt"
	"github.com/tsmoreland/go-data-examples/mongodb_example/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewMongoRepository(collection *mongo.Collection, ctx context.Context) (Repository, error) {
	if collection == nil {
		return nil, fmt.Errorf("invalid collection")
	}
	if ctx == nil {
		return nil, fmt.Errorf("invalid context")
	}

	return &MongoRepository{collection: collection, ctx: ctx}, nil
}

func (r *MongoRepository) AddPerson(person model.Person) (ObjectId, error) {
	result, err := r.collection.InsertOne(r.ctx, person)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}
func (r *MongoRepository) AddPeople(people []model.Person) ([]ObjectId, error) {
	//this fails (won't compile) need a means of converting people to bson.D (ordered) or bson.M (unordered)

	result, err := r.collection.InsertMany(r.ctx, people)
	if err != nil {
		return nil, err
	}
	var ids []ObjectId
	for _, id := range result.InsertedIDs {
		ids = append(ids, id)
	}
	return ids, nil
}

func (r *MongoRepository) FindById(id ObjectId) (*model.Person, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *MongoRepository) FindByFirstName(firstName string) ([]model.Person, error) {
	return []model.Person{}, fmt.Errorf("not implemented")
}
func (r *MongoRepository) FindByLastName(lastName string) ([]model.Person, error) {
	return []model.Person{}, fmt.Errorf("not implemented")
}

func (r *MongoRepository) UpdateById(id ObjectId, updatedPerson model.Person) error {
	return fmt.Errorf("not implemented")
}
func (r *MongoRepository) UpdateByFirstName(name string, updatedPerson model.Person) error {
	return fmt.Errorf("not implemented")
}
func (r *MongoRepository) UpdateByLastName(name string, updatedPerson model.Person) error {
	return fmt.Errorf("not implemented")
}

func (r *MongoRepository) DeleteById(id ObjectId) error {
	return fmt.Errorf("not implemented")
}
func (r *MongoRepository) DeleteByFirstName(name string) error {
	return fmt.Errorf("not implemented")
}
func (r *MongoRepository) DeleteByLastName(name string) error {
	return fmt.Errorf("not implemented")
}
