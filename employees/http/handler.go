package http

import (
	"github.com/gorilla/mux"

	"algogrit.com/emp-server/employees/service"
)

type EmployeeHandler struct {
	v1Svc service.EmployeeService
	*mux.Router
}

func (h *EmployeeHandler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/v1/employees", h.IndexV1).Methods("GET")
	r.HandleFunc("/v1/employees", h.CreateV1).Methods("POST")
}

func New(v1Svc service.EmployeeService) EmployeeHandler {
	h := EmployeeHandler{v1Svc: v1Svc}

	h.SetupRoutes(mux.NewRouter())

	return h
}
