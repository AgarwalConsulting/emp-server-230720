package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"algogrit.com/emp-server/entities"
)

func (h *EmployeeHandler) IndexV1(w http.ResponseWriter, r *http.Request) {
	emps, err := h.v1Svc.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emps)
}

func (h *EmployeeHandler) CreateV1(w http.ResponseWriter, r *http.Request) {
	var newEmp entities.Employee
	err := json.NewDecoder(r.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	validatorIns := validator.New()
	err = validatorIns.Struct(newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := h.v1Svc.Create(newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdEmp)
}
