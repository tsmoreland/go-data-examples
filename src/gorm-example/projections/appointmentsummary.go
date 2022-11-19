package projections

import "time"

type AppointmentSummary struct {
	FirstName string
	LastName  string
	Title     string
	StartTime time.Time
}
