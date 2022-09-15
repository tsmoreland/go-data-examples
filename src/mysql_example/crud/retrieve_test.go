package crud

import (
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"testing"
)

func Test_FindByIdShouldCallQueryWhenDatabaseAvailabile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	rows := sqlmock.
		NewRows([]string{"id", "first_name", "last_name"}).
		AddRow("1", "John", "Smith")

	mock.
		ExpectQuery("SELECT id, first_name, last_name FROM people WHERE").
		WillReturnRows(rows)

	p, err := FindById(db, 1)

	if err != nil {
		t.Fatal(err)
	}

	if p == nil {
		t.Fatalf("Query not executed")
	}
}
