package entities

import "github.com/jinzhu/gorm"

type Calendar struct {
	gorm.Model
	Name         string `sql:"index:idx_calendar_name""`
	EmployeeID   uint   `sql:"column:employee_id,index:idx_calendar_user_id"`
	Appointments []*Appointment
}

func (e Calendar) TableName() string {
	return "calendars"
}

func CreateCalendarTable(db *gorm.DB) {
	db.
		DropTableIfExists(&Calendar{}).
		CreateTable(&Calendar{}).
		Model(&Calendar{}).
		AddForeignKey("employee_id", "employees(id)", "CASCADE", "CASCADE")
}
