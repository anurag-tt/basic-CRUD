package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := Router()

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/employee/{id}", handler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/employees", handler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/add", handler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/edit/{id}", handler).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete/{id}", handler).Methods("DELETE", "OPTIONS")

	return router
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
