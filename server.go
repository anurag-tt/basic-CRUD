package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/anurag-tt/basic-CRUD/model"
	"github.com/anurag-tt/basic-CRUD/postgre"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
)

var (
	rEmployee postgre.Repository
	err       error
)

func main() {

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	rEmployee = postgre.InitDB()
	if err != nil {
		log.Fatalf("Error DB connection")
	}
	// fmt.Printf("%v", *rEmployee)

	r := Router()

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/employee/{id}", GetEmployee).Methods("GET")
	router.HandleFunc("/api/employees", GetAllEmployees).Methods("GET")
	router.HandleFunc("/api/add", CreateEmployee).Methods("POST")
	router.HandleFunc("/api/edit/{id}", UpdateEmployee).Methods("PUT")
	router.HandleFunc("/api/delete/{id}", DeleteEmployee).Methods("DELETE")

	return router
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	var emp model.EmployeeInfo
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		log.Errorf("[CreateEmployee] [Unable to decode the request body. ] [ERR] Error: %v\n", err)
	}

	insertID, err := rEmployee.CreateEmployee(emp)
	if err != nil {
		log.Errorf("[CreateEmployee] [Unable to Create an employee.] [ERR] Error: %v\n", err)
	}
	res := model.Response{
		ID:      insertID,
		Message: "Emplyee added successfully",
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Errorf("[CreateEmployee] [Encoding] [ERR] Error\n")
		return
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Errorf("[Employee] [Unable to convert the string into int] [ERR] Error: %v\n", err)
	}

	emp, err := rEmployee.GetEmployeeById(int64(id))
	if err != nil {
		log.Errorf("[Employee] [Unable to get employee.] [ERR] Error: %v\n", err)
	}

	res := model.Response{}
	if emp.Updated_At.IsZero() {
		res.ID = 0
		res.Message = "No Employee Found"

		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Errorf("[Employee] [Encoding] [ERR] Error: %v\n", err)
			return
		}
	} else {
		err = json.NewEncoder(w).Encode(emp)
		if err != nil {
			log.Errorf("[Employee] [Encoding] [ERR] Error: %v\n", err)
			return
		}
	}
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {

	emps, err := rEmployee.GetAllEmployees()
	if err != nil {
		log.Errorf("[Employee] [Unable to get all employees.] [ERR] Error: %v\n", err)
	}
	err = json.NewEncoder(w).Encode(emps)
	if err != nil {
		log.Errorf("[GetAllEmployees] [Encoding] [ERR] Error: %+v\n", err)
		return
	}
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Errorf("[Employee] [Unable to convert the string into int] [ERR] Error: %v\n", err)
	}

	var emp model.EmployeeInfo
	err = json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		log.Errorf("[UpdateEmployee] [Unable to decode the request body. ] [ERR] Error: %v\n", err)
	}

	updatedRows, err := rEmployee.UpdateEmployeeById(int64(id), emp)
	if err != nil {
		log.Errorf("[Employee] [Unable to update employee.] [ERR] Error: %v\n", err)
	}
	var msg string
	if updatedRows != 0 {
		msg = fmt.Sprintf("Employee updated successfully. Total rows/record affected %v", updatedRows)
	} else {
		msg = fmt.Sprintf("No rows affected. Employee with ID:%v may not exists.", id)
	}

	res := model.Response{
		ID:      int64(id),
		Message: msg,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Errorf("[UpdateEmployee] [Encoding] [ERR] Error\n")
		return
	}
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows, err := rEmployee.DeleteEmployeeById(int64(id))
	if err != nil {
		log.Errorf("[DeleteEmployee] [deletedRows] [ERR] Error\n")
	}
	var msg string
	if deletedRows != 0 {
		msg = fmt.Sprintf("Employee Deleted successfully. Total rows/record affected %v", deletedRows)
	} else {
		msg = fmt.Sprintf("No rows affected. Employee with ID:%v may not exists.", id)
	}

	res := model.Response{
		ID:      int64(id),
		Message: msg,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Errorf("[DeleteEmployee] [Encoding] [ERR] Error\n")
		return
	}
}
