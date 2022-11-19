package examples

import (
	"github.com/jinzhu/gorm"
	"github.com/tsmoreland/go-data-examples/src/gormexample/entities"
)

func FirstByLastNameOrInvalid(db *gorm.DB, lastName string) *entities.Employee {
	var employee entities.Employee
	db.
		Debug().
		Where("last_name", lastName).
		Attrs("id", 0).
		FirstOrInit(&employee)
	return &employee
}
