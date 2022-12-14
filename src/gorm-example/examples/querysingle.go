package examples

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/tsmoreland/go-data-examples/src/gormexample/entities"
	"github.com/tsmoreland/go-data-examples/src/gormexample/projections"
)

func FirstEmployeeIncludingCalendar(db *gorm.DB) {
	firstEmployee := entities.Employee{}
	firstEmployeeCalendar := entities.Calendar{}
	db.First(&firstEmployee).Related(&firstEmployeeCalendar)

	fmt.Println(firstEmployee) // e.Calendar will be nil, First does not load children by default
	fmt.Println(firstEmployeeCalendar)
}

// FirstEmployeeOrInit will initialize the entity if it does not already exist, but it will not save it to the database
func FirstEmployeeOrInit(db *gorm.DB, e *entities.Employee, constraint *entities.Employee) {
	db.Debug().FirstOrInit(e, constraint)
}

// FirstEmployeeOrCreate will create the entity if it does not already exist
func FirstEmployeeOrCreate(db *gorm.DB, e *entities.Employee, constraint *entities.Employee) {
	db.Debug().FirstOrInit(e, constraint)
}

func LastEmployee(db *gorm.DB) *entities.Employee {
	var e entities.Employee
	db.Debug().Last(&e)
	return &e
}

func LastEmployeeIncludingAppointments(db *gorm.DB) *entities.Employee {
	var e entities.Employee
	db.
		Debug().
		Preloads("Calendar.Appointments").
		Last(&e)
	return &e
}

func NextAppointmentForEmployee(db *gorm.DB, firstName string, lastName string) (*projections.AppointmentSummary, error) {
	var summaries []projections.AppointmentSummary

	db.
		Debug().
		Model(&entities.Employee{}).
		Joins("inner join calendars on calendars.employee_id").
		Joins("inner join appointments on appointments.calendar_id").
		Where("employees.first_name = ? and employees.last_name = ?", firstName, lastName).
		Limit(1).
		Select("employees.first_name, employees.last_name, calendars.name, appointments.title, appointments.start_time").
		Scan(&summaries)

	if len(summaries) > 0 {
		return &summaries[0], nil
	} else {
		return nil, fmt.Errorf("employee not found")
	}
}

func NextAppointmentForEmployeeUsingRows(db *gorm.DB, firstName string, lastName string) (*projections.AppointmentSummary, error) {

	rows, err := db.
		Debug().
		Model(&entities.Employee{}).
		Joins("inner join calendars on calendars.employee_id").
		Joins("inner join appointments on appointments.calendar_id").
		Where("employees.first_name = ? and employees.last_name = ?", firstName, lastName).
		Select("employees.first_name, employees.last_name, calendars.name, appointments.title, appointments.start_time").
		Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var summary projections.AppointmentSummary

		rows.Scan(&summary.FirstName, &summary.LastName, &summary.CalendarName, &summary.Title, &summary.StartTime)
		return &summary, nil
	}

	return nil, fmt.Errorf("employee not found")
}

func GetTotalAppointmentLengthByCalendarName(db *gorm.DB) (*map[string]int, error) {
	rows, err := db.
		Debug().
		Model(&entities.Appointment{}).
		Select("calendar_id, name, sum(length) as total_length").
		Group("calendar_id").
		Rows()
	if err != nil {
		return nil, err
	}

	countByName := make(map[string]int)
	for rows.Next() {
		var id int
		var length int
		var name string

		rows.Scan(&id, &name, &length)
		countByName[name] += length
	}

	return &countByName, nil
}

func GetTotalAppointmentLengthByCalendarNameHavingId(db *gorm.DB, calendarId int) (*map[string]int, error) {
	rows, err := db.
		Debug().
		Model(&entities.Appointment{}).
		Select("calendar_id, name, sum(length) as total_length").
		Group("calendar_id").
		Having("calendar_id = ?", calendarId).
		Rows()
	if err != nil {
		return nil, err
	}

	countByName := make(map[string]int)
	for rows.Next() {
		var id int
		var length int
		var name string

		rows.Scan(&id, &name, &length)
		countByName[name] += length
	}

	return &countByName, nil
}
