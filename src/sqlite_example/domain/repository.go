package domain

type Repository interface {
	Migrate() error
	Close() error

	Employees() EmployeeRepository
	Departments() EmployeeRepository
}

type EmployeeRepository interface {
	Create(employee Employee) (*Employee, error)
	FindById(id int) (*Employee, error)
	FindAll(pageNumber int, pageSize int) ([]Employee, error)
	Update(employee Employee) error
	Delete(employee Employee) error
}
type DepartmentRepository interface {
	Create(department Department) (*Department, error)
	FindById(id int) (*Department, error)
	FindAll(pageNumber int, pageSize int) ([]Department, error)
	Update(department Department) error
	Delete(department Department) error
}
