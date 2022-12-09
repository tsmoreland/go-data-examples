package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/tsmoreland/go-data-examples/mysql_example/configuration"
	"github.com/tsmoreland/go-data-examples/mysql_example/crud"
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
		DBName: config.DatabaseName(),
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

	// for use with transactions
	//ctx := context.Background()
	fmt.Println("Connected")

	if err := crud.DeleteAll(db); err != nil {
		panic(err)
	}
	if err := crud.SeedData(db); err != nil {
		panic(err)
	}
}
