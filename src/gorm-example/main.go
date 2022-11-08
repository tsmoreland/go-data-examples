package main

import (
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
	VerifyConnectionOrPanic(db)

	CreateTables(db)

}

func VerifyConnectionOrPanic(gormDB *gorm.DB) {
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

func CreateTables(db *gorm.DB) {
	db.CreateTable(&Employee{})
}

type Employee struct {
	ID          uint
	FirstName   string
	LastName    string
	JobCategory int32
}
