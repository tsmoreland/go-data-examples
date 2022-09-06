package crud

import "database/sql"

func UpdateFirstName(db *sql.DB, id int, firstname string) (int64, error) {
	result, err := db.Exec("UPDATE person SET first_name = ? WHERE id = ?", firstname, id)
	if err != nil {
		return 0, err
	}
	affectedId, err := result.RowsAffected()
	if err != nil {
		return 0, err
	} else {
		return affectedId, nil
	}
}
