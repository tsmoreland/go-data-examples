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

func GetAllEmployeesMatchingConstraint(db *gorm.DB, constraint *entities.Employee) []entities.Employee {
	var employees []entities.Employee
	db.Debug().Find(&employees, constraint)
	return employees
}

func GetAllEmployeesMatchingConstraints(db *gorm.DB, constraints map[string]interface{}) []entities.Employee {
	var employees []entities.Employee
	db.Debug().Find(&employees, constraints)
	return employees
}

func GetAllEmployeesMatchingInlineQuery(db *gorm.DB, where ...interface{}) []entities.Employee {
	var employees []entities.Employee
	db.Debug().Find(&employees, where)
	return employees
}

func DemoGetAll(db *gorm.DB) {
	all := GetAllEmployees(db)
	entities.PrintNames(all)

	all = GetAllEmployeesMatchingConstraint(db, &entities.Employee{LastName: "Wayne"})
	entities.PrintNames(all)

	all = GetAllEmployeesMatchingConstraints(db, map[string]interface{}{
		"last_name": "Prince",
	})
	entities.PrintNames(all)

	all = GetAllEmployeesMatchingInlineQuery(db, "last_name = ?", "Kent")
	entities.PrintNames(all)
}
