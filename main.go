package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var schema = `
 CREATE TABLE IF NOT EXISTS EMPLOYEE_INFO(
	ID  SERIAL PRIMARY KEY,
	NAME           TEXT      NOT NULL,
	DEPARTMENT      TEXT      NOT NULL,
	MANAGER_NAME   TEXT,
	JOINING_DATE   DATE
 );
 `

type employee_info struct {
	ID           int64  `db:"id"`
	Name         string `db:"name"`
	Department   string `db:"department"`
	Manager_Name string `db:"manager_name"`
	Joining_Date string `db:"joining_date"`
}

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/employee/{id}", GetEmployee).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/employees", GetAllEmployees).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/add", CreateEmployee).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/edit/{id}", UpdateEmployee).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete/{id}", DeleteEmployee).Methods("DELETE", "OPTIONS")

	return router
}

// DB Connection function
func createConnection() *sqlx.DB {

	connectionStr := "postgres://%s:%s@%s:%d/%s?sslmode=%s"
	psqlconn := fmt.Sprintf(connectionStr, os.Getenv("user"), os.Getenv("password"), "localhost", 5432, "employee", "disable")
	db, err := sqlx.Connect("postgres", psqlconn)
	CheckError(err)

	err = db.Ping()
	CheckError(err)

	db.MustExec(schema)

	fmt.Println("Successfully connected to DB!")
	return db

}

/*
	Route Functions
*/
func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	var emp employee_info
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := insertEmployee(emp)
	res := response{
		ID:      insertID,
		Message: "Emplyee added successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	emp, err := getEmployee(int64(id))
	if err != nil {
		log.Fatalf("Unable to get employee. %v", err)
	}

	json.NewEncoder(w).Encode(emp)
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {

	emps, err := getAllEmployees()
	if err != nil {
		log.Fatalf("Unable to get all employees. %v", err)
	}
	json.NewEncoder(w).Encode(emps)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var emp employee_info
	err = json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateEmployee(int64(id), emp)
	msg := fmt.Sprintf("Employee updated successfully. Total rows/record affected %v", updatedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteEmployee(int64(id))
	var msg string
	if deletedRows != 0 {
		msg = fmt.Sprintf("Employee deleted successfully. Total rows/record affected %v", deletedRows)
	} else {
		msg = fmt.Sprintf("Employee Do not exists!")
	}
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

/*
	Middleware Functions
*/
func insertEmployee(emp employee_info) int64 {

	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO employee_info (name, department, manager_name, joining_date) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, emp.Name, emp.Department, emp.Manager_Name, emp.Joining_Date).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}

func getEmployee(id int64) (employee_info, error) {
	db := createConnection()
	defer db.Close()

	var emp employee_info

	sqlStatement := `SELECT * FROM employee_info WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Manager_Name, &emp.Joining_Date)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return emp, nil
	case nil:
		return emp, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return emp, err
}

func getAllEmployees() ([]employee_info, error) {
	db := createConnection()
	defer db.Close()

	var emps []employee_info

	sqlStatement := `SELECT * FROM employee_info`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var emp employee_info
		err = rows.Scan(&emp.ID, &emp.Name, &emp.Department, &emp.Manager_Name, &emp.Joining_Date)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		emps = append(emps, emp)
	}

	return emps, err
}

func updateEmployee(id int64, emp employee_info) int64 {

	db := createConnection()
	defer db.Close()
	// sqlStatement := `UPDATE employee_info SET name=$2, department=$3, manager_name=$4, joining_date=$5 WHERE id=$1`

	// res, err := db.Exec(sqlStatement, id, emp.Name, emp.Department, emp.Manager_Name, emp.Joining_Date)
	// if err != nil {
	// 	log.Fatalf("Unable to execute the query. %v", err)
	// }

	tx := db.MustBegin()
	if emp.Name != "" {
		tx.MustExec(`UPDATE employee_info SET name=$2 WHERE id=$1`, id, emp.Name)
	}
	if emp.Department != "" {
		tx.MustExec(`UPDATE employee_info SET department=$2 WHERE id=$1`, id, emp.Department)
	}
	if emp.Manager_Name != "" {
		tx.MustExec(`UPDATE employee_info SET manager_name=$2 WHERE id=$1`, id, emp.Manager_Name)
	}
	if emp.Joining_Date != "" {
		tx.MustExec(`UPDATE employee_info SET joining_date=$2 WHERE id=$1`, id, emp.Joining_Date)
	}
	err := tx.Commit()
	// rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	// Issue: Not able to find affected rows.
	fmt.Printf("Total rows/record affected %v", 1)
	return 1
}

func deleteEmployee(id int64) int64 {

	db := createConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM employee_info WHERE id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}
