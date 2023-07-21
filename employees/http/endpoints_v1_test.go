package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/service"
	"algogrit.com/emp-server/entities"
)

func TestCreateV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)
	sut := empHTTP.New(mockSvc)

	expectedEmp := entities.Employee{Name: "Gaurav", Department: "LnD", ProjectID: 10001}

	mockSvc.EXPECT().Create(expectedEmp).Return(&expectedEmp, nil)

	jsonBody := `{"name": "Gaurav", "speciality": "LnD", "project": 10001}` // Type: string
	reqBody := strings.NewReader(jsonBody)

	req := httptest.NewRequest("POST", "/v1/employees", reqBody)
	resRec := httptest.NewRecorder()

	sut.CreateV1(resRec, req)
	// sut.ServeHTTP(resRec, req)

	assert.Equal(t, http.StatusOK, resRec.Code)

	var actualEmp entities.Employee
	json.NewDecoder(resRec.Body).Decode(&actualEmp)

	assert.Equal(t, expectedEmp, actualEmp)
}

func FuzzCreateV1(f *testing.F) {
	jsonBody := `{"name": "Gaurav", "speciality": "LnD", "project` // Type: string

	f.Add(jsonBody)

	f.Fuzz(func(t *testing.T, jsonBody string) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockSvc := service.NewMockEmployeeService(ctrl)
		sut := empHTTP.New(mockSvc)

		// expectedEmp := entities.Employee{Name: "Gaurav", Department: "LnD", ProjectID: 10001}

		// mockSvc.EXPECT().Create(expectedEmp).Return(&expectedEmp, nil)

		reqBody := strings.NewReader(jsonBody)

		req := httptest.NewRequest("POST", "/v1/employees", reqBody)
		resRec := httptest.NewRecorder()

		sut.CreateV1(resRec, req)
		// sut.ServeHTTP(resRec, req)

		assert.Equal(t, http.StatusBadRequest, resRec.Code)

		// var actualEmp entities.Employee
		// json.NewDecoder(resRec.Body).Decode(&actualEmp)

		// assert.Equal(t, expectedEmp, actualEmp)
	})
}
