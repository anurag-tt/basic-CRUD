package postgre

import "github.com/anurag-tt/basic-CRUD/model"

type Repository interface {
	GetEmployeeById(id int64) (emp model.EmployeeInfo, err error)
	GetAllEmployees() (emps []model.EmployeeInfo, err error)
	CreateEmployee(emp model.EmployeeInfo) (id int64, err error)
	UpdateEmployeeById(id int64, emp model.EmployeeInfo) (updatedRows int64, err error)
	DeleteEmployeeById(id int64) (deletedRows int64, err error)
}
