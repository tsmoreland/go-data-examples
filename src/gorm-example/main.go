package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	db, err := gorm.Open("postgres", "user=admin password=p@55w0rd!, dbname=gormExample, sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Print(err)
		}
	}()
	verifyConnectionOrPanic(db)
	resetTables(db)

	demoCrud(db)

}

func demoCrud(db *gorm.DB) {
	employee := Employee{
		FirstName:   "John",
		LastName:    "Smith",
		JobCategory: 1,
	}

	db.Create(&employee)
	printEmployee(employee)

	first := Employee{}

	db.First(&first)
	printEmployee(first)

	last := Employee{}
	db.Last(&last)
	printEmployee(last)

	update := Employee{LastName: "Smith"}
	db.Where(&update).First(&update)
	printEmployee(update)

	update.FirstName = "Jim"
	db.Save(&update)

	toRemove := Employee{LastName: "Smith"}
	db.Where(&toRemove).Delete(&toRemove)
}

func verifyConnectionOrPanic(gormDB *gorm.DB) {
	db := gormDB.DB()
	defer func() {
		if err := db.Close(); err != nil {
			log.Print(err)
		}
	}()

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	log.Print("Connected to database.")
}

func resetTables(db *gorm.DB) {
	db.DropTableIfExists(&Employee{})
	db.CreateTable(&Employee{})
}

func printEmployee(employee Employee) {
	fmt.Printf("%d: %s %s (%d)", employee.ID, employee.FirstName, employee.LastName, employee.JobCategory)
}

type Employee struct {
	ID          uint
	FirstName   string
	LastName    string
	JobCategory int32
}
