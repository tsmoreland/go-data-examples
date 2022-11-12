package entities

import "github.com/jinzhu/gorm"

type Calendar struct {
	gorm.Model
	Name       string
	EmployeeID uint `sql:"column:'employee_id'"`
}

func (e Calendar) TableName() string {
	return "calendars"
}

func CreateCalendarTable(db *gorm.DB) {
	db.
		DropTable(&Calendar{}).
		CreateTable(&Calendar{}).
		Model(&Calendar{}).
		AddForeignKey("employee_id", "employees(id)", "CASCADE", "CASCADE")
}
