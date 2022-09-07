package crud

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"testing"
)

func Test_AddShouldInsertIntoDatabaseWhenDatabaseAvailable(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	mock.
		ExpectExec("INSERT INTO People").
		WillReturnResult(sqlmock.NewResult(42, 1))

	id, err := Add(db, "John", "Smith")
	if err != nil {
		t.Fatal(err)
	}

	if id != 42 {
		t.Fatalf("Id %v does not match expected value %v", id, 42)
	}
}

func Test_AddShouldReturnErrorWhenInsertFails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	mock.
		ExpectExec("INSERT INTO People").
		WillReturnError(fmt.Errorf("expected failure"))

	_, err = Add(db, "John", "Smith")
	if err == nil {
		t.Fatal("error is nil")
	}
}
