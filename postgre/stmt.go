package postgre

import (
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/common/log"
)

func (rEmployee *EmployeeDBRepo) prepareStatements() {
	rEmployee.statement = make(map[string]*sqlx.Stmt)
	rEmployee.prepareGetEmployeeById()
	rEmployee.prepareGetAllEmployees()
	rEmployee.prepareCreateEmployee()
	rEmployee.prepareUpdateEmployeeById()
	rEmployee.prepareDeleteEmployeeById()
}

func (rEmployee *EmployeeDBRepo) prepareGetEmployeeById() {
	var (
		db              *sqlx.DB
		err             error
		getEmployeeById *sqlx.Stmt
	)

	db = rEmployee.db

	getEmployeeById, err = db.Preparex(queryGetEmployeeById)
	if err != nil {
		log.Errorf("[Employee] [prepareGetEmployeeById] [ERR] Preparing Statement getEmployeeById: %+v\n", err)
		return
	}

	rEmployee.statement[caseGetEmployeeById] = getEmployeeById
}

func (rEmployee *EmployeeDBRepo) prepareGetAllEmployees() {
	var (
		db              *sqlx.DB
		err             error
		getAllEmployees *sqlx.Stmt
	)

	db = rEmployee.db

	getAllEmployees, err = db.Preparex(queryGetAllEmployees)
	if err != nil {
		log.Errorf("[Employee] [prepareGetAllEmployees] [ERR] Preparing Statement getAllEmployees: %+v\n", err)
		return
	}

	rEmployee.statement[caseGetAllEmployees] = getAllEmployees
}

func (rEmployee *EmployeeDBRepo) prepareCreateEmployee() {
	var (
		db             *sqlx.DB
		err            error
		createEmployee *sqlx.Stmt
	)

	db = rEmployee.db

	createEmployee, err = db.Preparex(queryCreateEmployee)
	if err != nil {
		log.Errorf("[Employee] [prepareCreateEmployee] [ERR] Preparing Statement createEmployee: %+v\n", err)
		return
	}

	rEmployee.statement[caseCreateEmployee] = createEmployee
}

func (rEmployee *EmployeeDBRepo) prepareUpdateEmployeeById() {
	var (
		db                 *sqlx.DB
		err                error
		updateEmployeeById *sqlx.Stmt
	)

	db = rEmployee.db

	updateEmployeeById, err = db.Preparex(queryUpdateEmployeeById)
	if err != nil {
		log.Errorf("[Employee] [prepareUpdateEmployeeById] [ERR] Preparing Statement updateEmployeeById: %+v\n", err)
		return
	}

	rEmployee.statement[caseUpdateEmployeeById] = updateEmployeeById
}

func (rEmployee *EmployeeDBRepo) prepareDeleteEmployeeById() {
	var (
		db                 *sqlx.DB
		err                error
		deleteEmployeeById *sqlx.Stmt
	)

	db = rEmployee.db

	deleteEmployeeById, err = db.Preparex(queryDeleteEmployeeById)
	if err != nil {
		log.Errorf("[Eployee] [prepareDeleteEmployeeById] [ERR] Preparing Statement deleteEmployeeById: %+v\n", err)
		return
	}

	rEmployee.statement[caseDeleteEmployeeById] = deleteEmployeeById
}
