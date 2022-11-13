package models

import "fmt"

const (
	JobCategoryEngineer = iota
	JobCategoryIT
	JobCategorySupport
	JobCategorySales
	JobCategoryManager
	JobCategorySuperHero
)

type Employee struct {
	id          uint
	firstName   string
	lastName    string
	jobCategory int32
}

type EmployeeEntity struct {
	ID          uint
	FirstName   string
	LastName    string
	JobCategory int32
}

func (e EmployeeEntity) TableName() string {
	return "employee_entities"
}

func NewEmployee(firstname string, lastname string, jobCategory int32) (*Employee, error) {
	if len(firstname) == 0 {
		return nil, fmt.Errorf("firstname cannot be empty")
	}
	if len(lastname) == 0 {
		return nil, fmt.Errorf("lastname cannot be empty")
	}

	if jobCategory < 0 || jobCategory >= JobCategoryManager {
		return nil, fmt.Errorf("invalid job category")
	}

	return &Employee{firstName: firstname, lastName: lastname, jobCategory: jobCategory}, nil
}

func NewEmployeeWithId(id uint, firstname string, lastname string, jobCategory int32) (*Employee, error) {
	employee, err := NewEmployee(firstname, lastname, jobCategory)
	if err != nil {
		return nil, err
	}
	employee.id = id
	return employee, nil
}

func LoadFromEntity(e EmployeeEntity) *Employee {
	return &Employee{
		id:          e.ID,
		firstName:   e.FirstName,
		lastName:    e.LastName,
		jobCategory: e.JobCategory,
	}
}

func (e Employee) Print() {
	fmt.Printf("%d: %s %s (%d)", e.id, e.firstName, e.lastName, e.jobCategory)
}

func (e Employee) ToEntity() *EmployeeEntity {
	return &EmployeeEntity{
		ID:          e.id,
		FirstName:   e.firstName,
		LastName:    e.lastName,
		JobCategory: e.jobCategory,
	}
}
