package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/tsmoreland/go-data-examples/src/gormexample/entities"
	"github.com/tsmoreland/go-data-examples/src/gormexample/infrastructure"
	"github.com/tsmoreland/go-data-examples/src/gormexample/shared"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres password=P@55w0rd! dbname=gormexample sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer shared.CloseWithErrorLogging(db)
	sqlDB := db.DB()
	defer shared.CloseWithErrorLogging(sqlDB)
	infrastructure.VerifyConnectionOrPanic(sqlDB)
	infrastructure.ResetTables(db)

	infrastructure.CrudDemo(db)

	entities.CreateTables(db)
	entities.SeedDb(db)

	firstEmployee := entities.Employee{}
	firstEmployeeCalendar := entities.Calendar{}
	db.First(&firstEmployee).Related(&firstEmployeeCalendar)

	fmt.Println(firstEmployee) // e.Calendar will be nil, First does not load children by default
	fmt.Println(firstEmployeeCalendar)
}
