package infrastructure

import (
	"database/sql"
	"log"
)

func verifyConnectionOrPanic(db *sql.DB) {
	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	log.Print("Connected to database.")
}
