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
