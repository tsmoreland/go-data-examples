package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Appointment struct {
	gorm.Model
	Title       string
	Description string
	StartTime   time.Time
	Length      uint
	CalendarID  uint
	Attendees   []Employee `gorm:"many2many:appointment_employee"`
}

func (e Appointment) TableName() string {
	return "appointments"
}

func CreateAppointmentsTable(db *gorm.DB) {
	db.
		DropTableIfExists(&Appointment{}).
		CreateTable(&Appointment{})
}
