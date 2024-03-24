package models
import (
	"fmt"
	//
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	employeeTable = "employee"
	employeeFields = "id, fullname, email, phone, city"
)

type employee struct {
	ID int `json:"id"`
	FullName string `json:"fullname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	City string `json:"city"`
}

func (e *employee) GetScanArgs() []interface{} {
	return []interface{}{
		&e.ID,
		&e.FullName,
		&e.Email,
		&e.Phone,
		&e.City,
	}
}

// 
type EmployeeModel struct {
	mysqlDB *sql.DB;
}
//
func (e *EmployeeModel) GetEmployeeByEmail(email string) (*employee, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE email = ? LIMIT 1", employeeFields, employeeTable);
	row := e.mysqlDB.QueryRow(query, email);
	emp := &employee{};
	err := row.Scan(emp.GetScanArgs()...);
	if err != nil {
		return nil, err;
	}
	return emp, nil;
}
//
func NewEmployee(db *sql.DB) *EmployeeModel {
	return &EmployeeModel{
		mysqlDB: db,
	}
}