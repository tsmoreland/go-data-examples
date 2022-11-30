package domain

import "fmt"

type Employee struct {
	Id           int64
	FirstName    string
	LastName     string
	IsDeveloper  bool
	DepartmentId int64
	Department   *Department
}

func (e Employee) String() string {
	if e.Department != nil {
		return fmt.Sprintf("%v: %v %v (%v) in %v", e.Id, e.FirstName, e.LastName, e.IsDeveloper, e.Department)
	} else {
		return fmt.Sprintf("%v: %v %v (%v)", e.Id, e.FirstName, e.LastName, e.IsDeveloper)
	}
}

type Department struct {
	Id        int64
	Name      string
	Employees []Employee
}

func (d Department) String() string {
	return fmt.Sprintf("(%v: %v)", d.Id, d.Name)
}
