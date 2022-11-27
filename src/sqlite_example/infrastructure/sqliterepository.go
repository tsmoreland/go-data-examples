package infrastructure

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tsmoreland/go-data-examplessqlite_example/domain"
	"log"
	"os"
)

type SqliteRepository struct {
	db *sql.DB
}

func NewSqliteRepository(filename string) (domain.Repository, error) {
	if err := os.Remove(filename); err != nil {
		log.Println(err)
	}

	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	r := &SqliteRepository{db: db}
	return r, nil
}

//goland:noinspection ALL
func (r SqliteRepository) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS People (
	  id INT AUTO_INCREMENT PRIMARY KEY,
	  first_name varchar(100),
	  last_name varchar(100)
	)`
	result, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	_ = result
	return nil
}

func (r SqliteRepository) Close() error {
	return r.db.Close()
}
