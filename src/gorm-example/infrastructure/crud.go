package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/tsmoreland/go-data-examples/src/gormexample/models"
)

func CrudDemo(db *gorm.DB) {
	employee := models.EmployeeEntity{
		FirstName:   "John",
		LastName:    "Smith",
		JobCategory: 1,
	}

	db.Create(&employee)
	models.LoadFromEntity(employee).Print()

	first := models.EmployeeEntity{}

	db.First(&first)
	models.LoadFromEntity(employee).Print()

	last := models.EmployeeEntity{}
	db.Last(&last)
	models.LoadFromEntity(employee).Print()

	update := models.EmployeeEntity{LastName: "Smith"}
	db.Where(&update).First(&update)
	models.LoadFromEntity(employee).Print()

	update.FirstName = "Jim"
	models.LoadFromEntity(update).Print()

	toRemove := models.EmployeeEntity{LastName: "Smith"}
	db.Where(&toRemove).Delete(&toRemove)
}
