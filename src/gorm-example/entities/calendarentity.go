package entities

import "github.com/jinzhu/gorm"

type CalendarEntity struct {
	gorm.Model
	EmployeeID uint
}

func (e CalendarEntity) TableName() string {
	return "calendars"
}
