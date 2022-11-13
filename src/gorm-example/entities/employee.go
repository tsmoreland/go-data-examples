package entities

import "github.com/jinzhu/gorm"

type Employee struct {
	ID          uint
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
		DropTableIfExists(&Employee{}).
		CreateTable(&Employee{}).
		Model(&Employee{}).AddIndex("idx_employee_job_category_name", "job_category_name")
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
