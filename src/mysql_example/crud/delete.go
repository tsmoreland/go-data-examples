package crud

import "database/sql"

func DeleteById(db *sql.DB, id int64) (int64, error) {
	result, err := db.Exec("DELETE FROM People WHERE id = ?", id)
	if err != nil {
		return 0, nil
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	} else {
		return affected, nil
	}
}
