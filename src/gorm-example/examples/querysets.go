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

func FindByLastNames(db *gorm.DB, lastNames ...string) []entities.Employee {
	var employees []entities.Employee

	for _, lastName := range lastNames {
		var matches []entities.Employee
		db.Debug().Where(map[string]interface{}{"last_name": lastName}).Find(&matches)

		for _, match := range matches {
			employees = append(employees, match)
		}
	}
	return employees
}
