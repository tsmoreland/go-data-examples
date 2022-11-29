package infrastructure

import (
	"github.com/tsmoreland/go-data-examplessqlite_example/domain"
	"github.com/tsmoreland/go-data-examplessqlite_example/shared"
)

type scanner interface {
	Scan(dest ...any) error
}

func readEmployee(rows scanner) (*domain.Employee, error) {
	var emp domain.Employee
	var isDeveloper int
	if err := rows.Scan(&emp.Id, &emp.FirstName, &emp.LastName, &isDeveloper, &emp.DepartmentId); err != nil {
		return nil, translate(err)
	}
	if isDeveloper != 0 {
		emp.IsDeveloper = true
	} else {
		emp.IsDeveloper = false
	}
	return &emp, nil
}

func readEmployeeIncludingDepartment(rows scanner, department *domain.Department) (*domain.Employee, error) {
	emp, err := readEmployee(rows)
	if err != nil {
		return nil, translate(err)
	}

	if emp.DepartmentId != department.Id {
		return nil, shared.ErrInvalidArgument
	}
	return emp, nil
}

func readDepartment(rows scanner) (*domain.Department, error) {
	var department domain.Department
	if err := rows.Scan(&department.Id, &department.Name); err != nil {
		return nil, translate(err)
	}
	return &department, nil
}
