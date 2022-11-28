package infrastructure

import (
	"database/sql"
	"errors"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tsmoreland/go-data-examplessqlite_example/domain"
	"github.com/tsmoreland/go-data-examplessqlite_example/shared"
	"log"
	"os"
)

type SqliteRepository struct {
	db *sql.DB
}

func NewSqliteRepository(filename string) (domain.Repository, error) {
	if err := os.Remove(filename); err != nil {
		log.Println(err)
	}

	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	r := &SqliteRepository{db: db}
	return r, nil
}

//goland:noinspection ALL
func (r *SqliteRepository) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS Department (
	   id INT AUTO_INCREMENT PRIMARY KEY,
	   name varchar(50)
	);
	`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS Employees (
		id INT AUTO_INCREMENT PRIMARY KEY,
	  	first_name varchar(100),
	  	last_name varchar(100),
	    is_developer int,
	    department_id int not null,
	    CONSTRAINT fk_departments
	    	FOREIGN KEY (department_id)
	    	REFERENCES departments(department_id)
	);`
	_, err = r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// SeedData adds initial data for the app to test with
func (r *SqliteRepository) SeedData() error {
	return shared.ErrNotImplemented
}

// Close closes the underlying connection to the database
func (r *SqliteRepository) Close() error {
	return r.db.Close()
}

// Employees returns an instance of EmployeeRepository
func (r *SqliteRepository) Employees() domain.EmployeeRepository {
	return r
}

// Departments returns an instance of DepartmentRepository
func (r *SqliteRepository) Departments() domain.DepartmentRepository {
	return r
}

func (r *SqliteRepository) CreateEmployee(employee domain.Employee) (*domain.Employee, error) {
	command := "INSERT INTO Employees (id, first_name, last_name, is_developer, department_id) VALUES (?, ?, ?, ?)"
	isDeveloper := 0
	if employee.IsDeveloper {
		isDeveloper = 1
	}

	res, err := r.db.Exec(command, employee.FirstName, employee.LastName, isDeveloper, employee.DepartmentId)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return nil, shared.ErrDuplicate
		}
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	employee.Id = id

	return &employee, nil
}
func (r *SqliteRepository) FindEmployeeById(id int) (*domain.Employee, error) {
	return nil, shared.ErrNotImplemented
}
func (r *SqliteRepository) FindAllEmployees(pageNumber int, pageSize int) ([]domain.Employee, error) {
	return nil, shared.ErrNotImplemented
}
func (r *SqliteRepository) UpdateEmployee(employee domain.Employee) error {
	return shared.ErrNotImplemented
}
func (r *SqliteRepository) DeleteEmployee(employee domain.Employee) error {
	return shared.ErrNotImplemented
}

func (r *SqliteRepository) CreateDepartment(department domain.Department) (*domain.Department, error) {
	command := "INSERT INTO Departments (name) VALUES (?)"

	res, err := r.db.Exec(command, department.Name)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return nil, shared.ErrDuplicate
		}
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	department.Id = id

	if department.Employees == nil {
		department.Employees = make([]domain.Employee, 0)
	}

	var employees []domain.Employee
	for _, employee := range department.Employees {

		// TODO: wrap this in a transaction - either all employees are added or none are
		employee.DepartmentId = department.Id
		newEmployee, err := r.CreateEmployee(employee)
		if err == nil {
			employees = append(employees, *newEmployee)
			continue
		}
		if err = r.DeleteDepartment(department); err != nil {
			return nil, err
		}
	}

	return &department, nil
}
func (r *SqliteRepository) FindDepartmentById(id int) (*domain.Department, error) {
	query := "SELECT * FROM Departments WHERE id = ?"

	row := r.db.QueryRow(query, id)
	var department domain.Department
	if err := row.Scan(&department.Id, &department.Name); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, shared.ErrNotFound
		}
		return nil, err
	}

	// TODO:
	// add a method that can be used to scan the rows for employees to be used here and by
	// find all employees
	query = "SELECT * from Employees WHERE department_id = ?"
	rows, err := r.db.Query(query, department.Id)
	if err != nil {
		return nil, err
	}
	defer shared.CloseWithErrorReporting(rows)

	var employees []domain.Employee
	for rows.Next() {
		var employee domain.Employee
		var isDeveloper int
		if err := rows.Scan(&employee.Id, &employee.FirstName, &employee.LastName, isDeveloper, &employee.DepartmentId); err != nil {
			return nil, err
		}
		if isDeveloper != 0 {
			employee.IsDeveloper = true
		} else {
			employee.IsDeveloper = false
		}
		employee.Department = &department
		employees = append(employees, employee)
	}

	department.Employees = employees
	return &department, nil
}
func (r *SqliteRepository) FindAllDepartments(pageNumber int, pageSize int, includeEmployees bool) ([]domain.Department, error) {
	return nil, shared.ErrNotImplemented
}
func (r *SqliteRepository) UpdateDepartment(department domain.Department) error {
	return shared.ErrNotImplemented
}
func (r *SqliteRepository) DeleteDepartment(department domain.Department) error {
	return shared.ErrNotImplemented
}
