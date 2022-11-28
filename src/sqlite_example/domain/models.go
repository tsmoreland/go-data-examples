package domain

type Employee struct {
	Id          int64
	FirstName   string
	LastName    string
	IsDeveloper bool
	Department  Department
}

type Department struct {
	Id        int64
	Name      string
	Employees []Employee
}
