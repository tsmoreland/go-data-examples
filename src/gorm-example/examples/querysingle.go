package examples

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/tsmoreland/go-data-examples/src/gormexample/entities"
)

func FirstEmployeeIncludingCalendar(db *gorm.DB) {
	firstEmployee := entities.Employee{}
	firstEmployeeCalendar := entities.Calendar{}
	db.First(&firstEmployee).Related(&firstEmployeeCalendar)

	fmt.Println(firstEmployee) // e.Calendar will be nil, First does not load children by default
	fmt.Println(firstEmployeeCalendar)
}

// FirstEmployeeOrInit will initialize the entity if it does not already exist, but it will not save it to the database
func FirstEmployeeOrInit(db *gorm.DB, e *entities.Employee, constraint *entities.Employee) {
	db.Debug().FirstOrInit(e, constraint)
}

// FirstEmployeeOrCreate will create the entity if it does not already exist
func FirstEmployeeOrCreate(db *gorm.DB, e *entities.Employee, constraint *entities.Employee) {
	db.Debug().FirstOrInit(e, constraint)
}

func LastEmployee(db *gorm.DB) *entities.Employee {
	var e entities.Employee
	db.Debug().Last(&e)
	return &e
}

func LastEmployeeIncludingAppointments(db *gorm.DB) *entities.Employee {
	var e entities.Employee
	db.
		Debug().
		Preloads("Calendar.Appointments").
		Last(&e)
	return &e
}
