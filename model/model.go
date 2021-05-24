package model

import (
	"time"

	"github.com/jackc/pgtype"
	"github.com/jmoiron/sqlx"
)

type EmployeeInfo struct {
	ID           int64       `json:"id"`
	Name         string      `json:"name"`
	Department   string      `json:"department"`
	Manager_Name string      `json:"manager_name,omitempty"`
	Joining_Date pgtype.Date `json:"joining_date,omitempty"`
	Updated_At   time.Time   `json:"updated_at"`
}

type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type EmployeeDBRepo struct {
	db        *sqlx.DB
	statement map[string]*sqlx.Stmt
}
