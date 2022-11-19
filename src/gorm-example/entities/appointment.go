package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Appointment represents a calendar entry.  The Recurring field should be commented out when first creating tables to
// demonstrate updating an existing table
type Appointment struct {
	gorm.Model
	Title       string
	Description string
	StartTime   time.Time
	Length      uint
	CalendarID  uint
	Recurring   bool
	Attendees   []*Employee `gorm:"many2many:appointment_employee"`
	Attachments []Attachment
}

func (e Appointment) TableName() string {
	return "appointments"
}

func CreateAppointmentsTable(db *gorm.DB) {
	db.
		CreateTable(&Appointment{})
}
