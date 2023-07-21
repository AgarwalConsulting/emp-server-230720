package repository_test

import (
	"sync"
	"testing"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/entities"
	"github.com/stretchr/testify/assert"
)

func TestConsistency(t *testing.T) {
	sut := repository.NewInMem()

	emps, err := sut.ListAll()

	assert.Nil(t, err)
	assert.Equal(t, 3, len(emps))

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			emp := entities.Employee{Name: "Gaurav"}

			createdEmp, err := sut.Save(emp)

			assert.Nil(t, err)
			assert.NotNil(t, createdEmp)

			emps, err := sut.ListAll()

			assert.Nil(t, err)
			assert.NotNil(t, emps)
		}()
	}

	wg.Wait()

	emps, err = sut.ListAll()

	assert.Nil(t, err)
	assert.Equal(t, 103, len(emps))
}
