package domain

type Repository interface {
	Migrate() error
	Close() error

	Employees() EmployeeRepository
	Departments() DepartmentsRepository
}

type EmployeeRepository interface {
	CreateEmployee(employee Employee) (*Employee, error)
	FindEmployeeById(id int) (*Employee, error)
	FindAllEmployees(pageNumber int, pageSize int) ([]Employee, error)
	UpdateEmployee(employee Employee) error
	DeleteEmployee(employee Employee) error
}
type DepartmentsRepository interface {
	CreateDepartment(department Department) (*Department, error)
	FindDepartmentById(id int) (*Department, error)
	FindAllDepartments(pageNumber int, pageSize int, includeEmployees bool) ([]Department, error)
	UpdateDepartment(department Department) error
	DeleteDepartment(department Department) error
}
