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

func (r *MongoRepository) AddPerson(person model.Person) (ObjectId, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *MongoRepository) AddPeople(people []model.Person) ([]ObjectId, error) {
	return []ObjectId{}, fmt.Errorf("not implemented")
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
