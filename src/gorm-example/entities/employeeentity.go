package entities

type EmployeeEntity struct {
	ID          uint
	FirstName   string `sql:"type:VARCHAR(100)"`
	LastName    string `sql:"size:100"`
	JobCategory int32
}

func (e EmployeeEntity) TableName() string {
	return "employees"
}
