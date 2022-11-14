package examples

import (
	"github.com/jinzhu/gorm"
	"github.com/tsmoreland/go-data-examples/src/gormexample/entities"
)

func GetAllEmployees(db *gorm.DB) []entities.Employee {
	var employees []entities.Employee
	db.Debug().Find(&employees)
	return employees
}

func GetAllEmployeesMatching(db *gorm.DB, constraint *entities.Employee) []entities.Employee {
	var employees []entities.Employee
	db.Debug().Find(&employees, constraint)
	return employees
}
