package infrastructure

import (
	"database/sql"
	"errors"
	"github.com/mattn/go-sqlite3"
	"github.com/tsmoreland/go-data-examplessqlite_example/shared"
)

func translate(err error) error {
	if err == nil {
		return nil
	}

	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
		return shared.ErrDuplicate
	}
	if errors.Is(err, sql.ErrNoRows) {
		return shared.ErrNotFound
	}

	return err
}
