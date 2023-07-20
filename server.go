package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"project"`
}

var employees = []Employee{
	{1, "Gaurav", "LnD", 10001},
	{2, "Thamim", "Cloud", 20001},
	{3, "Shagun", "SRE", 20002},
}

func EmployeesIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, r *http.Request) {
	var newEmp Employee
	err := json.NewDecoder(r.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmp.ID = len(employees) + 1

	employees = append(employees, newEmp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmp)
}

func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		EmployeeCreateHandler(w, req)
	} else {
		EmployeesIndexHandler(w, req)
	}
}

func main() {
	r := mux.NewRouter()
	// r := http.NewServeMux()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	// r.HandleFunc("/employees", EmployeesHandler)
	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	http.ListenAndServe("localhost:8000", r)
}
