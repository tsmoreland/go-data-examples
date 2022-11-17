package examples

import (
	"github.com/jinzhu/gorm"
	"github.com/tsmoreland/go-data-examples/src/gormexample/entities"
	"github.com/tsmoreland/go-data-examples/src/gormexample/models"
)

func GetPagedEmployees(db *gorm.DB, pageNumber int, pageSize int) []entities.Employee {
	skip := (pageNumber - 1) * pageSize
	take := pageSize

	var employees []entities.Employee
	db.
		Debug().
		Order("last_name DESC").
		Offset(skip).
		Limit(take).
		Preloads("Calendar.Appointments").
		Find(&employees)
	return employees
}

func GetPagedProjections(db *gorm.DB, pageNumber int, pageSize int) []entities.Employee {
	skip := (pageNumber - 1) * pageSize
	take := pageSize
	var employees []entities.Employee
	db.
		Debug().
		Order("last_name DESC").
		Offset(skip).
		Limit(take).
		Select([]string{"first_name", "last_name"}).
		Find(&employees)

	return employees
}

func FindByName(db *gorm.DB, firstName string, lastName string) []entities.Employee {
	var employees []entities.Employee
	db.
		Debug().
		Where("first_name = ? AND last_name = ?", firstName, lastName).
		Find(&employees)
	return employees
}

func FindByNameWithCalendar(db *gorm.DB, firstName string, lastName string) []entities.Employee {
	var employees []entities.Employee
	db.
		Debug().
		Where("first_name = ? AND last_name = ?", firstName, lastName).
		Preloads("Calendar.Appointments").
		Find(&employees)
	return employees
}

func FindByLastNames(db *gorm.DB, lastNames ...string) []entities.Employee {
	var employees []entities.Employee

	query := db.Debug()
	first := true
	for _, lastName := range lastNames {
		if first {
			query = query.Where(map[string]interface{}{"last_name": lastName})
			first = false
		} else {
			query = query.Or(map[string]interface{}{"last_name": lastName})
		}
	}
	query.Find(&employees)
	return employees
}

func FindByFirstName(db *gorm.DB, firstNames ...string) []entities.Employee {
	var employees []entities.Employee

	db.Debug().Where("first_name in (?)", firstNames).Find(&employees)

	return employees
}

func FindAllNonVillains(db *gorm.DB) []entities.Employee {
	var employees []entities.Employee
	db.Debug().Not("jobcategory_id = ?", models.JobCategoryVillain).Find(&employees)
	return employees
}
