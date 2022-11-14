package main

import (
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

	infrastructure.CreateTables(db)
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

}
