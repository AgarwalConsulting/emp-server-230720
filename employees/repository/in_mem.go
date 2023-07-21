package repository

import (
	"sync"

	"algogrit.com/emp-server/entities"
)

type inmemRepo struct {
	employees []entities.Employee
	mut       sync.RWMutex
}

func (repo *inmemRepo) ListAll() ([]entities.Employee, error) {
	repo.mut.RLock()
	defer repo.mut.RUnlock()

	return repo.employees, nil
}

func (repo *inmemRepo) Save(newEmp entities.Employee) (*entities.Employee, error) {
	repo.mut.Lock()
	defer repo.mut.Unlock()

	newEmp.ID = len(repo.employees) + 1
	repo.employees = append(repo.employees, newEmp)

	return &newEmp, nil
}

func NewInMem() EmployeeRepository {
	var employees = []entities.Employee{
		{1, "Gaurav", "LnD", 10001},
		{2, "Thamim", "Cloud", 20001},
		{3, "Shagun", "SRE", 20002},
	}

	return &inmemRepo{employees: employees}
}
