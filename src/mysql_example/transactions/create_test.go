package transactions

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"testing"
)

func Test_AddIfNotPresentShouldReturnErrorIfValueExists(t *testing.T) {
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
		NewRows([]string{"id"}).
		AddRow(1)
	mock.ExpectBegin()
	mock.
		ExpectQuery("SELECT id from people where").
		WillReturnRows(rows)

	ctx := context.Background()
	_, err = AddIfNotPresent(db, ctx, "John", "Smith")

	if err == nil {
		t.Fatal(err)
	}
}

func Test_AddIfNotPresentShouldNotReturnNewIdAndNilError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Println(err)
		}
	}()

	ctx := context.Background()
	mock.ExpectBegin()

	rows := sqlmock.NewRows([]string{"id"})
	mock.
		ExpectQuery("SELECT id from people where").
		WillReturnRows(rows)
	mock.
		ExpectExec("INSERT INTO People \\(first_name, last_name\\)").
		WillReturnResult(sqlmock.NewResult(42, 1))
	mock.ExpectCommit()

	_, err = AddIfNotPresent(db, ctx, "John", "Smith")
	if err != nil {
		t.Fatal(err)
	}
}
