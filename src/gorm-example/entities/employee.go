package entities

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Employee struct {
	gorm.Model
	FirstName   string          `sql:"type:VARCHAR(100);not null" gorm:"column:first_name"`
	LastName    string          `sql:"size:100;not null"`
	JobCategory JobCategoryLink `gorm:"embedded"`
	Calendar    Calendar
}

type JobCategoryLink struct {
	JobCategoryID   uint   `sql:"not null;DEFAULT:'1'"`
	JobCategoryName string `sql:"size:100;not null;DEFAULT:'Engineer'"`
}

func CreateEmployeeTable(db *gorm.DB) {
	db.
		CreateTable(&Employee{})
}

func (e *Employee) TableName() string {
	return "employees"
}

func (e *Employee) BeforeUpdate() error {
	return nil
}
func (e *Employee) AfterUpdate() error {
	return nil
}

func (e *Employee) AddAppointment(a *Appointment) {
	e.Calendar.Appointments = append(e.Calendar.Appointments, a)
}

func PrintNames(employees []Employee) {
	fmt.Printf("(%v) Employees:", len(employees))
	for _, e := range employees {
		fmt.Printf("%v %v\n", e.FirstName, e.LastName)
	}
}
