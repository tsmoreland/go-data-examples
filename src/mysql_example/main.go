package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/tsmoreland/go-data-examples/mysql_example/configuration"
	"log"
)

var db *sql.DB

func main() {

	config, err := configuration.NewBuilder().
		AddJsonFile("settings.json").
		AddUserSecrets().
		AddEnvironment().
		Build()
	if err != nil {
		log.Fatal(err)
	}

	dsn := mysql.Config{
		User:   config.Username(),
		Passwd: config.Password(),
		Net:    "tcp",
		Addr:   config.Address(),
		DBName: "example",
	}
	db, err = sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")
}
