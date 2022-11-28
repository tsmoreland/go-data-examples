package domain

type Employee struct {
	Id          uint
	FirstName   string
	LastName    string
	IsDeveloper bool
	Department  Department
}

type Department struct {
	Id        uint
	Name      string
	Employees []Employee
}
