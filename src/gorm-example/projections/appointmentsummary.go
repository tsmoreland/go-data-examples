package projections

import "time"

type AppointmentSummary struct {
	FirstName    string
	LastName     string
	CalendarName string `gorm:"column:name"`
	Title        string
	StartTime    time.Time
}
