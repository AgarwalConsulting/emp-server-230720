package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/entities"
)

var empRepo = repository.NewInMem()

func EmployeesIndexHandler(w http.ResponseWriter, r *http.Request) {
	emps, err := empRepo.ListAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emps)
}

func EmployeeCreateHandler(w http.ResponseWriter, r *http.Request) {
	var newEmp entities.Employee
	err := json.NewDecoder(r.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := empRepo.Save(newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdEmp)
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
