package entities

import "github.com/jinzhu/gorm"

type Appointment struct {
	gorm.Model
	Name string
}

func (e Appointment) TableName() string {
	return "appointments"
}

func CreateAppointmentsTable(db *gorm.DB) {
	db.
		DropTable(&Appointment{}).
		CreateTable(&Appointment{})
}
