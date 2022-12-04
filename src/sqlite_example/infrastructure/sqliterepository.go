package infrastructure

import (
	"context"
	"database/sql"
	"errors"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tsmoreland/go-data-examplessqlite_example/domain"
	"github.com/tsmoreland/go-data-examplessqlite_example/shared"
	"log"
	"math"
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
	CREATE TABLE IF NOT EXISTS Departments (
	   id INTEGER PRIMARY KEY AUTOINCREMENT,
	   name varchar(50)
	);
	`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	query = `
	CREATE TABLE IF NOT EXISTS Employees (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
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
	departments := []domain.Department{
		{Id: 1, Name: "Finance"},
		{Id: 2, Name: "Legal"},
		{Id: 3, Name: "Marketing"},
		{Id: 4, Name: "HR"},
		{Id: 5, Name: "IT"},
		{Id: 6, Name: "R&D"},
	}
	for _, entity := range departments {
		if _, err := r.CreateDepartment(entity); err != nil {
			return err
		}
	}

	employees := []domain.Employee{
		{Id: 1, FirstName: "John", LastName: "Smith", IsDeveloper: false, DepartmentId: 1},
		{Id: 2, FirstName: "Jessica", LastName: "Jones", IsDeveloper: true, DepartmentId: 5},
		{Id: 3, FirstName: "Tony", LastName: "Stark", IsDeveloper: true, DepartmentId: 6},
		{Id: 4, FirstName: "Edward", LastName: "Nigma", IsDeveloper: false, DepartmentId: 1},
		{Id: 5, FirstName: "Brenda", LastName: "Moore", IsDeveloper: false, DepartmentId: 4},
		{Id: 6, FirstName: "Bruce", LastName: "Wayne", IsDeveloper: true, DepartmentId: 6},
		{Id: 7, FirstName: "Bruce", LastName: "Banner", IsDeveloper: true, DepartmentId: 6},
		{Id: 8, FirstName: "Harley", LastName: "Quinn", IsDeveloper: false, DepartmentId: 4},
		{Id: 9, FirstName: "Victor", LastName: "Fries", IsDeveloper: false, DepartmentId: 6},
		{Id: 10, FirstName: "Harvey", LastName: "Dent", IsDeveloper: false, DepartmentId: 2},
	}

	for _, entity := range employees {
		if _, err := r.CreateEmployee(entity); err != nil {
			return err
		}
	}

	return nil
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

type dbExecutor interface {
	Exec(query string, args ...any) (sql.Result, error)
}

func (r *SqliteRepository) CreateEmployee(employee domain.Employee) (*domain.Employee, error) {
	return createEmployeeUsingDbExecutor(r.db, employee)
}

func createEmployeeUsingDbExecutor(executor dbExecutor, employee domain.Employee) (*domain.Employee, error) {
	command := "INSERT INTO Employees (first_name, last_name, is_developer, department_id) VALUES (?, ?, ?, ?)"
	isDeveloper := 0
	if employee.IsDeveloper {
		isDeveloper = 1
	}

	res, err := executor.Exec(command, employee.FirstName, employee.LastName, isDeveloper, employee.DepartmentId)
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

func (r *SqliteRepository) FindEmployeeById(id int64) (*domain.Employee, error) {
	_ = id
	return nil, shared.ErrNotImplemented
}
func (r *SqliteRepository) FindAllEmployees(pageNumber int, pageSize int) ([]domain.Employee, error) {
	if pageNumber < 1 || pageSize < 1 {
		return nil, shared.ErrInvalidArgument
	}

	query := `
select
  *
from
  Employees
  inner join Departments on Employees.department_id = Departments.id
order by 
  last_name
limit 
  ? 
offset 
  ?
`
	rows, err := r.db.Query(query, pageSize, (pageNumber-1)*pageSize)
	if err != nil {
		return nil, err
	}
	defer shared.CloseWithErrorReporting(rows)

	var employees []domain.Employee
	for rows.Next() {
		employee, err := readEmployeeWithEmbeddedDepartment(rows)
		if err != nil {
			return nil, err
		}
		employees = append(employees, *employee)
	}
	return employees, nil
}

func (r *SqliteRepository) FindAllEmployeesInDepartment(department *domain.Department, pageNumber int, pageSize int) ([]domain.Employee, error) {
	var query string

	if pageNumber < 1 || pageSize < 1 {
		return nil, shared.ErrInvalidArgument
	}

	if pageSize == math.MaxInt {
		query = "SELECT * from Employees WHERE department_id = ?"
	}

	rows, err := r.db.Query(query, department.Id)
	if err != nil {
		return nil, err
	}
	defer shared.CloseWithErrorReporting(rows)

	var employees []domain.Employee
	for rows.Next() {
		employee, err := readEmployeeIncludingDepartment(rows, department)
		if err != nil {
			return nil, err
		}
		employees = append(employees, *employee)
	}
	return employees, nil
}

func (r *SqliteRepository) FindAllEmployeesWithDepartmentId(departmentId int64, pageNumber int, pageSize int) ([]domain.Employee, error) {
	department, err := r.FindDepartmentById(departmentId, false)
	if err != nil {
		return nil, translate(err)
	}
	employees, err := r.FindAllEmployeesInDepartment(department, pageNumber, pageSize)
	if err != nil {
		return nil, translate(err)
	}

	return employees, nil
}

func (r *SqliteRepository) UpsertEmployee(employee domain.Employee) (*domain.Employee, error) {
	if employee.Id == 0 {
		return r.CreateEmployee(employee)
	}

	var isDeveloper int
	if employee.IsDeveloper {
		isDeveloper = 1
	} else {
		isDeveloper = 0
	}

	command := `
UPDATE Employees SET 
  first_name = ?, 
  last_name = ?, 
  is_developer = ?, 
  department_id = ?
WHERE 
  id = ?
`

	res, err := r.db.Exec(command, employee.FirstName, employee.LastName, isDeveloper, employee.DepartmentId, employee.Id)
	if err != nil {
		return nil, translate(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, translate(err)
	}
	if rowsAffected == 0 {
		return nil, shared.ErrNotFound
	}

	return r.FindEmployeeById(employee.Id)
}
func (r *SqliteRepository) DeleteEmployee(employee domain.Employee) error {
	res, err := r.db.Exec("DELETE FROM Employees WHERE id = ?", employee.Id)
	if err != nil {
		return translate(err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return translate(err)
	}

	if rowsAffected == 0 {
		return shared.ErrDeleteFailed
	}
	return nil
}

func (r *SqliteRepository) CreateDepartment(department domain.Department) (*domain.Department, error) {
	command := "INSERT INTO Departments (name) VALUES (?)"

	res, err := r.db.Exec(command, department.Name)
	if err := translate(err); err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	department.Id = id

	if department.Employees == nil {
		department.Employees = make([]domain.Employee, 0)
	}

	tx, err := r.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, translate(err)
	}

	var employees []domain.Employee
	for _, employee := range department.Employees {

		employee.DepartmentId = department.Id
		newEmployee, err := createEmployeeUsingDbExecutor(tx, employee)
		if err == nil {
			employees = append(employees, *newEmployee)
			continue
		}
		if err = r.DeleteDepartment(department); err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, translate(err)
	}

	return &department, nil
}
func (r *SqliteRepository) FindDepartmentById(id int64, includeEmployees bool) (*domain.Department, error) {
	query := "SELECT * FROM Departments WHERE id = ?"

	row := r.db.QueryRow(query, id)
	department, err := readDepartment(row)
	if err != nil {
		return nil, translate(err)
	}

	if !includeEmployees {
		department.Employees = make([]domain.Employee, 0)
		return department, nil
	}

	employees, err := r.FindAllEmployeesInDepartment(department, 1, math.MaxInt)
	if err != nil {
		return nil, translate(err)
	}
	department.Employees = employees
	return department, nil
}

func (r *SqliteRepository) FindAllDepartments(pageNumber int, pageSize int, includeEmployees bool) ([]domain.Department, error) {
	if pageNumber < 1 || pageSize < 1 {
		return nil, shared.ErrInvalidArgument
	}

	query := "SELECT * FROM Departments"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, translate(err)
	}
	defer shared.CloseWithErrorReporting(rows)

	var departments []domain.Department
	for rows.Next() {
		department, err := readDepartment(rows)
		if err != nil {
			return nil, translate(err)
		}
		departments = append(departments, *department)
		// first pass; better solution is to get all employees and add them to a map of departments to employees - since
		// every employee has to have an apartment, means 2 overall queries rather than 1 + n where n
		// is the number of departments
		if includeEmployees {
			employees, err := r.FindAllEmployeesInDepartment(department, 1, math.MaxInt)
			if err != nil {
				return nil, translate(err)
			}
			department.Employees = employees
		}
	}

	return departments, nil
}
func (r *SqliteRepository) UpsertDepartment(department domain.Department) (*domain.Department, error) {
	if department.Id == 0 {
		return r.CreateDepartment(department)
	}

	return nil, shared.ErrNotImplemented
}
func (r *SqliteRepository) DeleteDepartment(department domain.Department) error {
	_ = department
	return shared.ErrNotImplemented
}
