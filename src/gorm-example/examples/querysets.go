package examples

import (
	"github.com/jinzhu/gorm"
	"github.com/tsmoreland/go-data-examples/src/gormexample/entities"
)

func FindByName(db *gorm.DB, firstName string, lastName string) []entities.Employee {
	var employees []entities.Employee
	db.
		Debug().
		Where("first_name = ? AND last_name = ?", firstName, lastName).
		Find(&employees)
	return employees
}
