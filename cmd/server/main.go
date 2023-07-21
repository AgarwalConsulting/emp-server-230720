package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	var empRepo = repository.NewInMem()
	var empSvc = service.NewV1(empRepo)
	var empHandler = empHTTP.New(empSvc)

	empHandler.SetupRoutes(r)

	http.ListenAndServe(":8000", r)
}
