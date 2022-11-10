package apis

import (
	"bytes"
	"employee_manage/controllers/department"
	. "employee_manage/tests/config"
	"employee_manage/tests/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDepartments(t *testing.T) {
	mocks.MockGetDepartments()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/departments/", nil)
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	success := response["success"].(bool)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, success)
}

func TestGetDepartmentByID(t *testing.T) {
	tests := []TestModel{
		{
			Name: "Test Get Department By ID",
			Type: "Success",
			Args: nil,
			Mock: mocks.MockGetDepartmentByID,
			Want: http.StatusOK,
		},
		{
			Name: "Test Fail Get Department By ID",
			Type: "Fail",
			Args: nil,
			Mock: mocks.MockFailGetDepartmentByID,
			Want: http.StatusBadRequest,
		},
		{
			Name: "Test Protect Role",
			Type: "Role",
			Args: nil,
			Mock: mocks.MockUser,
			Want: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		tt.Mock()
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tt.Args)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/departments/3", bytes.NewBuffer(body))
		if tt.Type == "Role" {
			req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
		} else {
			req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
		}
		router.ServeHTTP(w, req)

		var response department.MessageResponse
		json.Unmarshal(w.Body.Bytes(), &response)

		if tt.Type == "Success" {
			assert.Equal(t, tt.Want, w.Code)
			assert.NotEmpty(t, response)
			assert.True(t, response.Success)
		} else if tt.Type == "Fail" {
			assert.Equal(t, tt.Want, w.Code)
			assert.NotEmpty(t, response)
			assert.False(t, response.Success)
		}
	}
}

func TestCreateDepartment(t *testing.T) {
	tests := []TestModel{
		{
			Name: "Test Root Create Department",
			Type: "Success",
			Args: department.NewDepartment{
				Name:           "Phòng X",
				DepartmentCode: "DX",
				Address:        "Tầng X",
				Status:         true,
			},
			Mock: mocks.MockCreateDepartment,
			Want: http.StatusCreated,
		},
		{
			Name: "Test Protect Role",
			Type: "Role",
			Args: nil,
			Mock: mocks.MockUser,
			Want: http.StatusForbidden,
		},
		{
			Name: "Test Validate",
			Type: "Validate",
			Args: department.NewDepartment{
				DepartmentCode: "DX",
				Address:        "Tầng X",
				Status:         true,
			},
			Mock: mocks.MockAdmin,
			Want: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		tt.Mock()
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tt.Args)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/departments/", bytes.NewBuffer(body))
		if tt.Type == "Role" {
			req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
		} else {
			req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
		}
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		var response department.MessageResponse
		json.Unmarshal(w.Body.Bytes(), &response)

		if tt.Type == "Success" {
			assert.Equal(t, tt.Want, w.Code)
			assert.NotEmpty(t, response)
			assert.True(t, response.Success)
		} else if tt.Type == "Fail" {
			assert.Equal(t, tt.Want, w.Code)
			assert.NotEmpty(t, response)
			assert.False(t, response.Success)
		} else {
			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)
			assert.Equal(t, tt.Want, w.Code)
			assert.NotEmpty(t, response)
			assert.False(t, response["success"].(bool))
		}
	}
}

func TestUpdateDepartmentByID(t *testing.T) {
	tests := []TestModel{
		{
			Name: "Test Root Update Department By ID",
			Type: "Success",
			Args: map[string]interface{}{
				"name":            "Phòng Y",
				"department_code": "DY",
				"address":         "Tầng Y",
				"status":          false,
			},
			Mock: mocks.MockUpdateDepartmentByID,
			Want: http.StatusOK,
		},
		{
			Name: "Test Protect Role",
			Type: "Role",
			Args: nil,
			Mock: mocks.MockUser,
			Want: http.StatusForbidden,
		},
		{
			Name: "Test Validate",
			Type: "Validate",
			Args: map[string]interface{}{
				"department_code": "DY",
				"address":         "Tầng Y",
				"status":          false,
			},
			Mock: mocks.MockAdmin,
			Want: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		tt.Mock()
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tt.Args)
		req, _ := http.NewRequest(http.MethodPut, "/api/v1/departments/4", bytes.NewBuffer(body))
		if tt.Type == "Role" {
			req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
		} else {
			req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
		}
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		var response department.MessageResponse
		json.Unmarshal(w.Body.Bytes(), &response)

		if tt.Type == "Success" {
			assert.Equal(t, tt.Want, w.Code)
			assert.NotEmpty(t, response)
			assert.True(t, response.Success)
		} else if tt.Type == "Fail" {
			assert.Equal(t, tt.Want, w.Code)
			assert.NotEmpty(t, response)
			assert.False(t, response.Success)
		} else {
			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)
			assert.Equal(t, tt.Want, w.Code)
			assert.NotEmpty(t, response)
			assert.False(t, response["success"].(bool))
		}
	}
}

func TestDeleteDepartmentByID(t *testing.T) {
	tests := []TestModel{
		{
			Name: "Test Root Delete Department By ID",
			Type: "Success",
			Args: nil,
			Mock: mocks.MockDeleteDepartmentByID,
			Want: http.StatusOK,
		},
		{
			Name: "Test Root Delete Fail Department By ID",
			Type: "Fail",
			Args: nil,
			Mock: mocks.MockFailDeleteDepartmentByID,
			Want: http.StatusNotFound,
		},
		{
			Name: "Test Protect Role",
			Type: "Role",
			Args: nil,
			Mock: mocks.MockUser,
			Want: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		tt.Mock()
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tt.Args)
		req, _ := http.NewRequest(http.MethodDelete, "/api/v1/departments/4", bytes.NewBuffer(body))
		if tt.Type == "Role" {
			req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
		} else {
			req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
		}
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		var response department.MessageResponse
		json.Unmarshal(w.Body.Bytes(), &response)

		if tt.Type == "Success" {
			assert.Equal(t, tt.Want, w.Code)
			assert.NotEmpty(t, response)
			assert.True(t, response.Success)
		} else if tt.Type == "Fail" {
			assert.Equal(t, tt.Want, w.Code)
			assert.NotEmpty(t, response)
			assert.False(t, response.Success)
		} else {
			var response map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &response)
			assert.Equal(t, tt.Want, w.Code)
			assert.NotEmpty(t, response)
			assert.False(t, response["success"].(bool))
		}
	}
}
