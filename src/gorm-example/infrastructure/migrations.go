package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/tsmoreland/go-data-examples/src/gormexample/models"
)

func ResetTables(db *gorm.DB) {
	db.DropTableIfExists(&models.EmployeeEntity{})
	db.CreateTable(&models.EmployeeEntity{})
}
