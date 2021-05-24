package postgre

import (
	"github.com/jmoiron/sqlx"
)

type (

	// EmployeeDBRepo struct of repo employee with postgre
	EmployeeDBRepo struct {
		db        *sqlx.DB
		statement map[string]*sqlx.Stmt
	}
)

const (
	//postgres://[user]:[pswd]@[host]:[port]/[dbName]?sslmode=[sslmode]
	connectionStr = `postgres://%s:%s@%s:%d/%s?sslmode=%s`
)
