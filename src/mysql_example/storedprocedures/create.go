package storedprocedures

import (
	"database/sql"
	"github.com/tsmoreland/go-data-examples/mysql_example/models"
	"log"
)

func AddIfNotPresent(db *sql.DB, firstname string, lastname string) ([]models.Person, error) {

	results, err := db.Query("CALL addPersonIfNotExist(?, ?)", firstname, lastname)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := results.Close(); err != nil {
			log.Println(err)
		}
	}()

	var people []models.Person
	for results.Next() {
		var person models.Person
		if err := results.Scan(&person.Id, &person.Firstname, &person.Lastname); err != nil {
			return nil, err
		}
		people = append(people, person)
		if err := results.Err(); err != nil {
			return nil, err
		}
	}
	return people, err
}
