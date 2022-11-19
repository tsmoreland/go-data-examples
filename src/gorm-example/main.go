package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/tsmoreland/go-data-examples/src/gormexample/entities"
	"github.com/tsmoreland/go-data-examples/src/gormexample/examples"
	"github.com/tsmoreland/go-data-examples/src/gormexample/infrastructure"
	"github.com/tsmoreland/go-data-examples/src/gormexample/models"
	"github.com/tsmoreland/go-data-examples/src/gormexample/shared"
)

func main() {
	models.IgnoreUnused()

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=P@55w0rd! dbname=gormexample sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer shared.CloseWithErrorLogging(db)
	sqlDB := db.DB()
	defer shared.CloseWithErrorLogging(sqlDB)
	infrastructure.VerifyConnectionOrPanic(sqlDB)
	infrastructure.ResetCrudTables(db)

	infrastructure.CrudDemo(db)

	// use false when demonstrating adding columns to existing db
	infrastructure.CreateTables(db, true)
	if err := infrastructure.SeedDb(db); err != nil {
		panic(err)
	}

	examples.FirstEmployeeIncludingCalendar(db)

	harley := entities.Employee{
		FirstName: "Harley",
		LastName:  "Quinn",
	}
	harleyCalendar := entities.Calendar{}
	examples.UpdateHarleyQuinn(db, &harley, &harleyCalendar)
	examples.BulkUpdateWhenSetValueIsCalculated(db)
	examples.BulkUpdateWhenSetValueIsCalculated(db)
	examples.BulkDelete(db)

	var martianManHunter entities.Employee
	examples.FirstEmployeeOrInit(db, &martianManHunter, &entities.Employee{FirstName: "J'onn", LastName: "J'onzz"})

	examples.DemoGetAll(db)
	employees := examples.FindByName(db, "Bruce", "Wayne")
	entities.PrintNames(employees)

	employees = examples.FindByNameWithCalendar(db, "Bruce", "Wayne")
	entities.PrintNames(employees)

	employees = examples.FindByLastNames(db, "Wayne", "Kent", "Prince")
	entities.PrintNames(employees)

	employees = examples.FindByFirstName(db, "Bruce", "Clark", "Diana")
	entities.PrintNames(employees)

	employees = examples.FindAllNonVillains(db)
	entities.PrintNames(employees)

	employees = examples.GetPagedEmployees(db, 2, 2)
	entities.PrintNames(employees)

	appointmentSummary, err := examples.NextAppointmentForEmployee(db, "Bruce", "Wayne")
	if err == nil {
		fmt.Printf("%v %v: %v - %v at %v",
			appointmentSummary.FirstName,
			appointmentSummary.LastName,
			appointmentSummary.CalendarName,
			appointmentSummary.Title,
			appointmentSummary.StartTime)
	}

	appointmentSummary, err = examples.NextAppointmentForEmployeeUsingRows(db, "Bruce", "Wayne")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v %v: %v - %v at %v\n",
		appointmentSummary.FirstName,
		appointmentSummary.LastName,
		appointmentSummary.CalendarName,
		appointmentSummary.Title,
		appointmentSummary.StartTime)

	countByName, err := examples.GetTotalAppointmentLengthByCalendarName(db)
	if err != nil {
		panic(err)
	}
	for key, value := range *countByName {
		fmt.Printf("%v: %v\n", key, value)
	}
}
