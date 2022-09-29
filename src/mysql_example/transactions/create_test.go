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
