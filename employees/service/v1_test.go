package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

func TestMain(m *testing.M) {
	// Setup
	m.Run()
	// Teardown
}

func TestV1Index(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockEmployeeRepository(ctrl)
	sut := service.NewV1(mockRepo)

	expectedEmps := []entities.Employee{
		{1, "Gaurav", "LnD", 1010},
	}

	mockRepo.EXPECT().ListAll().Return(expectedEmps, nil)

	actualEmps, err := sut.Index()

	assert.Nil(t, err)
	assert.NotNil(t, actualEmps)

	assert.Equal(t, expectedEmps, actualEmps)
}
