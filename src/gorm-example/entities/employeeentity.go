package entities

import "github.com/jinzhu/gorm"

type EmployeeEntity struct {
	ID          uint
	FirstName   string          `sql:"type:VARCHAR(100);not null" gorm:"column:first_name"`
	LastName    string          `sql:"size:100;not null"`
	JobCategory JobCategoryLink `gorm:"embedded"`
}

type JobCategoryLink struct {
	JobCategoryID   uint   `sql:"not null;DEFAULT:'1'"`
	JobCategoryName string `sql:"size:100;not null;DEFAULT:'Engineer'"`
}

func (e EmployeeEntity) TableName() string {
	return "employees"
}

func CreateEmployeeTable(db *gorm.DB) {
	db.
		CreateTable(&EmployeeEntity{}).
		Model(&EmployeeEntity{}).AddIndex("idx_employee_job_category_name", "job_category_name")
}
