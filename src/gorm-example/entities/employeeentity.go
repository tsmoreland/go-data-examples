package entities

type EmployeeEntity struct {
	ID          uint
	FirstName   string `sql:"type:VARCHAR(100);not null"`
	LastName    string `sql:"size:100;not null"`
	JobCategory int32  `sql:DEFAULT:1`
}

func (e EmployeeEntity) TableName() string {
	return "employees"
}
