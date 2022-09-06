package crud

import "database/sql"

func Add(db *sql.DB, firstname string, lastname string) (int64, error) {
	result, err := db.Exec("INSERT INTO People (first_name, last_name,) VALUES (?, ?)", firstname, lastname)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	} else {
		return id, nil
	}
}
