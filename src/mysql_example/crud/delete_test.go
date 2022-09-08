package crud

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"testing"
)

func Test_DeleteShouldExecDeleteQueryWhenDatabaseAvailable(t *testing.T) {
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
		ExpectExec("DELETE FROM People WHERE id").
		WillReturnResult(sqlmock.NewResult(0, 1))

	affected, err := DeleteById(db, 42)
	if err != nil {
		t.Fatal(err)
	}

	if affected != 1 {
		t.Fatalf("affected rows %v did not match expected 1", affected)
	}
}

func Test_DeleteShouldReturnZeroAndErrWhenExecFails(t *testing.T) {
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
		ExpectExec("DELETE FROM People WHERE id").
		WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("error")))

	affected, err := DeleteById(db, 42)
	if err == nil {
		t.Fatal(err)
	}

	if affected != 0 {
		t.Fatalf("affected row count is not zero")
	}

}
