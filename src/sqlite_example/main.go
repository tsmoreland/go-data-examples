package main

import (
	"github.com/tsmoreland/go-data-examplessqlite_example/infrastructure"
	"github.com/tsmoreland/go-data-examplessqlite_example/shared"
)

func main() {
	r, err := infrastructure.NewSqliteRepository("example.db")
	if err != nil {
		panic(err)
	}
	defer shared.CloseWithErrorReporting(r)

	if err := r.Migrate(); err != nil {
		panic(err)
	}

}
