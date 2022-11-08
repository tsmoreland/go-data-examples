package models

import "fmt"

const (
	JobCategoryEngineer = iota
	JobCategoryIT
	JobCategorySupport
	JobCategorySales
	JobCategoryManager
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

func NewEmployee(id uint, firstname string, lastname string, jobCategory int32) (*Employee, error) {
	employee, err := NewEmployee(firstname, lastname, jobCategory)
	if err != nil {
		return nil, err
	}
	employee.id = id
	return employee, nil
}

func (e Employee) ToEntity() *EmployeeEntity {
	return &EmployeeEntity{
		ID: e.id,
		FirstName: e.firstName,
		LastName: e.lastName,
		JobCategory: e.jobCategory
	}
}