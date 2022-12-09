package crud

import "database/sql"

func SeedData(db *sql.DB) error {

	if _, err := Add(db, "Bruce", "Wayne"); err != nil {
		return err
	}
	if _, err := Add(db, "Clark", "Kent"); err != nil {
		return err
	}
	if _, err := Add(db, "Diana", "Prince"); err != nil {
		return err
	}
	if _, err := Add(db, "Hal", "Jordan"); err != nil {
		return err
	}
	if _, err := Add(db, "Barry", "Allen"); err != nil {
		return err
	}
	return nil
}
