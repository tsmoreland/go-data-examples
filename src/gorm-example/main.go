package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
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
}
