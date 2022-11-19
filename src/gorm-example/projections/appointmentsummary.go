package projections

import "time"

type AppointmentSummary struct {
	FirstName string
	LastName  string
	Title     String
	StartTime time.Time
}
