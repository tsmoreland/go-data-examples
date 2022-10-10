package infrastructure

import "github.com/tsmoreland/go-data-examples/mongodb_example/model"

type ObjectId interface{}

type Repository interface {
	AddPerson(person model.Person) (ObjectId, error)
	AddPeople(people []model.Person) ([]ObjectId, error)

	FindById(id ObjectId) (*model.Person, error)
	FindByFirstName(firstName string) ([]model.Person, error)
	FindByLastName(lastName string) ([]model.Person, error)

	UpdateById(id ObjectId, updatedPerson model.Person) error
	UpdateByFirstName(name string, updatedPerson model.Person) error
	UpdateByLastName(name string, updatedPerson model.Person) error

	DeleteById(id ObjectId) error
	DeleteByFirstName(name string) error
	DeleteByLastName(name string) error
}
