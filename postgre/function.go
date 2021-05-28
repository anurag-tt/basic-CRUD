package postgre

import (
	"database/sql"
	"fmt"

	"github.com/anurag-tt/basic-CRUD/model"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/common/log"
)

func (rEmployee *EmployeeDBRepo) GetEmployeeById(id int64) (emp model.EmployeeInfo, err error) {
	var stmt *sqlx.Stmt

	if stmt = rEmployee.statement[caseGetEmployeeById]; stmt == nil {
		log.Errorf("[Employee] [Get Statement] [ERR] Error: %v\n", err)
		return
	}

	err = stmt.QueryRow(id).Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Manager_Name, &emp.Joining_Date, &emp.Updated_At)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return emp, nil
	case nil:
		return emp, nil
	default:
		log.Errorf("[Employee] [Get Employee Query] [ERR] Error: %v\n", err)
	}

	return
}

func (rEmployee *EmployeeDBRepo) GetAllEmployees() (emps []model.EmployeeInfo, err error) {
	var stmt *sqlx.Stmt

	if stmt = rEmployee.statement[caseGetAllEmployees]; stmt == nil {
		log.Errorf("[Employee] [Get All Statement] [ERR] Error: %v\n", err)
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Errorf("[Employee] [Get All Employee Query] [ERR] Error: %v\n", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var emp model.EmployeeInfo
		err = rows.Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Manager_Name, &emp.Joining_Date, &emp.Updated_At)
		if err != nil {
			log.Errorf("[Employee] [Get All Employee Scan] [ERR] Error: %v\n", err)
			return
		}
		emps = append(emps, emp)
	}
	return
}

func (rEmployee *EmployeeDBRepo) CreateEmployee(emp model.EmployeeInfo) (id int64, err error) {
	var stmt *sqlx.Stmt

	if stmt = rEmployee.statement[caseCreateEmployee]; stmt == nil {
		log.Errorf("[Employee] [Create Employee Statement] [ERR] Error: %v\n", err)
		return
	}

	err = stmt.QueryRow(&emp.Name, &emp.Department, &emp.Manager_Name, &emp.Joining_Date).Scan(&id)
	if err != nil {
		log.Errorf("[Employee] [Create Employee Query] [ERR] Error: %v\n", err)
		return
	}
	return id, nil
}

func (rEmployee *EmployeeDBRepo) UpdateEmployeeById(id int64, emp model.EmployeeInfo) (updatedRows int64, err error) {
	var stmt *sqlx.Stmt

	if stmt = rEmployee.statement[caseUpdateEmployeeById]; stmt == nil {
		log.Errorf("[Employee] [Update Employee Statement] [ERR] Error: %v\n", err)
		return
	}

	rows, err := stmt.Exec(id, &emp.Name, &emp.Department, &emp.Manager_Name, &emp.Joining_Date)
	if err != nil {
		log.Errorf("[Employee] [Update Employee Query] [ERR] Error: %v\n", err)
		return
	}
	return rows.RowsAffected()
}

func (rEmployee *EmployeeDBRepo) DeleteEmployeeById(id int64) (deletedRows int64, err error) {
	var stmt *sqlx.Stmt

	if stmt = rEmployee.statement[caseDeleteEmployeeById]; stmt == nil {
		log.Errorf("[Employee] [Delete Employee Statement] [ERR] Error: %v\n", err)
		return
	}

	rows, err := stmt.Exec(id)
	if err != nil {
		log.Errorf("[Employee] [Delete Employee Query] [ERR] Error: %v\n", err)
		return
	}
	return rows.RowsAffected()
}
