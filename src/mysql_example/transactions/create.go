package transactions

import (
	"context"
	"database/sql"
	"log"
)

func AddIfNotPresent(db *sql.DB, ctx context.Context, firstname string, lastname string) (int64, error) {

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Println(err)
		}
	}()

	var existingId int64
	if err = tx.QueryRowContext(ctx, "SELECT id from people where first_name = ? and last_name = ?", firstname,
		lastname).Scan(&existingId); err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
	}

	if existingId > 0 {
		if err = tx.Rollback(); err != nil {
			return 0, err
		}
		// person exists
		return existingId, nil
	}

	result, err := tx.ExecContext(ctx, "INSERT INTO People (first_name, last_name,) VALUES (?, ?)", firstname, lastname)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	} else {
		return id, nil
	}
}
