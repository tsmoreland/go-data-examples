package domain

type Repository interface {
	Migrate() error
	SeedData() error
	Close() error

	Employees() EmployeeRepository
	Departments() DepartmentRepository
}

type EmployeeRepository interface {
	CreateEmployee(employee Employee) (*Employee, error)
	FindEmployeeById(id int) (*Employee, error)
	FindAllEmployees(pageNumber int, pageSize int) ([]Employee, error)
	UpsertEmployee(employee Employee) (*Employee, error)
	DeleteEmployee(employee Employee) error
}
type DepartmentRepository interface {
	CreateDepartment(department Department) (*Department, error)
	FindDepartmentById(id int, includeEmployees bool) (*Department, error)
	FindAllDepartments(pageNumber int, pageSize int, includeEmployees bool) ([]Department, error)
	UpsertDepartment(department Department) (*Department, error)
	DeleteDepartment(department Department) error
}
