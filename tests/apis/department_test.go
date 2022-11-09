package apis

import (
	"bytes"
	"employee_manage/controllers/department"
	"employee_manage/controllers/user"
	. "employee_manage/tests/config"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDepartments(t *testing.T) {
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
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/departments/1", nil)
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}

var departmentID string

func TestCreateDepartment(t *testing.T) {
	w := httptest.NewRecorder()

	newDepartment := department.NewDepartment{
		Name:           "Phòng X",
		DepartmentCode: "DX",
		Address:        "Tầng X",
		Status:         true,
	}
	body, _ := json.Marshal(newDepartment)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/departments/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	data, _ := response.Data.(map[string]interface{})
	departmentID = strconv.Itoa(int(data["id"].(float64)))

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}

func TestValidateCreateDepartment(t *testing.T) {
	w := httptest.NewRecorder()

	newDepartment := department.NewDepartment{
		DepartmentCode: "DX",
		Address:        "Tầng X",
		Status:         true,
	}
	body, _ := json.Marshal(newDepartment)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/departments/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	success := response["success"].(bool)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, response)
	assert.False(t, success)
}

func TestProtectRoleCreateDepartment(t *testing.T) {
	w := httptest.NewRecorder()

	newDepartment := department.NewDepartment{
		Name:           "Phòng X",
		DepartmentCode: "DX",
		Address:        "Tầng X",
		Status:         true,
	}
	body, _ := json.Marshal(newDepartment)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/departments/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+JWTAccountManager.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.NotEmpty(t, response)
	assert.False(t, response.Success)
}

func TestUpdateDepartmentByID(t *testing.T) {
	w := httptest.NewRecorder()

	dataUpdate := map[string]interface{}{
		"name":            "Phòng Y",
		"department_code": "DY",
		"address":         "Tầng Y",
		"status":          false,
	}

	body, _ := json.Marshal(dataUpdate)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/departments/3", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}

func TestDeleteDepartmentByID(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/departments/"+(departmentID), nil)
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}
