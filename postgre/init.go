package postgre

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	jsoniter "github.com/json-iterator/go"
	"github.com/prometheus/common/log"
)

var Jsoni = jsoniter.ConfigFastest

func InitDB() Repository {
	var err error
	NewDb := &EmployeeDBRepo{}

	connStr := fmt.Sprintf(connectionStr, os.Getenv("user"), os.Getenv("password"), "localhost", 5432, "employee", "disable")
	NewDb.db, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	NewDb.prepareStatements()
	return NewDb
}
