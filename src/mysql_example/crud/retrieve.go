package crud

import (
	"database/sql"
	"fmt"
	"github.com/tsmoreland/go-data-examples/mysql_example/models"
)

func FindById(db *sql.DB, id int64) ([]models.Person, error) {
	if db == nil {
		return nil, fmt.Errorf("db cannot be nil")
	}

	var people []models.Person
	results, err := db.Query("SELECT id, first_name, last_name FROM people WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	for results.Next() {
		var person models.Person
		if err := results.Scan(&person.Id, &person.Firstname, &person.Lastname); err != nil {
			return nil, fmt.Errorf("parse row err %v %v", id, err)
		}
		people = append(people, person)

		if err := results.Err(); err != nil {
			return nil, fmt.Errorf("result err %v %v", id, err)
		}
	}
	return people, nil
}
