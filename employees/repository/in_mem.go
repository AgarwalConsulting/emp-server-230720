package repository

import "algogrit.com/emp-server/entities"

type inmemRepo struct {
	employees []entities.Employee
}

func (repo *inmemRepo) ListAll() ([]entities.Employee, error) {
	return repo.employees, nil
}

func (repo *inmemRepo) Save(newEmp entities.Employee) (*entities.Employee, error) {
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

	return &inmemRepo{employees}
}
