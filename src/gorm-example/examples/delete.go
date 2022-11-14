package examples

import (
	"github.com/jinzhu/gorm"
	"github.com/tsmoreland/go-data-examples/src/gormexample/entities"
)

func BulkDelete(db *gorm.DB) {
	db.
		Debug().
		Where("last_name LIKE ?", "%mite").
		Delete(&entities.Employee{})
}
