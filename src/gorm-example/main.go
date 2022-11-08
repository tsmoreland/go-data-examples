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

	dbase := db.DB()
	defer func() {
		if err := dbase.Close(); err != nil {
			log.Print(err)
		}
	}()

	if err := dbase.Ping(); err != nil {
		panic(err.Error())
	}

	log.Print("Connected to database.")
}
