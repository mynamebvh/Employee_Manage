package test

import (
	"bytes"
	"employee_manage/controllers/user"
	. "employee_manage/test/config"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRootGetUserByID(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/3", nil)
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}

var userID string

func TestRootCreateUser(t *testing.T) {
	w := httptest.NewRecorder()
	birthday, _ := time.Parse("2006-01-02 15:04:05.000 -0700", "2020-01-02 03:04:05.000 +0000")

	newUser := user.NewUser{
		FullName:     "Bui Viet Hoang",
		Password:     "hoangdz",
		Phone:        "0966150922",
		Email:        "mynamebvh19@gmail.com",
		Gender:       true,
		Address:      "Hà Nội",
		DepartmentID: 1,
		RoleID:       3,
		Birthday:     birthday,
	}
	body, _ := json.Marshal(newUser)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	data, _ := response.Data.(map[string]interface{})
	userID = strconv.Itoa(int(data["id"].(float64)))

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}

func TestRootUpdateUserByID(t *testing.T) {
	w := httptest.NewRecorder()

	userUpdate := map[string]interface{}{
		"full_name":     "test update",
		"email":         "mynamebvh3000@gmail.com",
		"phone":         "0979850933",
		"address":       "Hà Nội",
		"department_id": 1,
		"gender":        true,
		"birthday":      "2020-01-02T15:04:05",
	}

	body, _ := json.Marshal(userUpdate)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/3", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}

func TestRootDeleteUserByID(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/users/"+(userID), nil)
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}
