package test

import (
	"bytes"
	"employee_manage/controllers/request"
	"employee_manage/controllers/user"
	. "employee_manage/test/config"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequests(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/requests/", nil)
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}

func TestGetRequestByID(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/requests/1", nil)
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}

func TestValidateRequest(t *testing.T) {
	w := httptest.NewRecorder()

	newRequest := request.NewRequest{
		Type: "take_leave",
	}
	body, _ := json.Marshal(newRequest)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/requests/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	success := response["success"].(bool)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, response)
	assert.False(t, success)
}

func TestAuthRequest(t *testing.T) {
	w := httptest.NewRecorder()

	newRequest := request.NewRequest{
		Type:    "take_leave",
		Content: "Em đưa vợ đi đẻ",
	}
	body, _ := json.Marshal(newRequest)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/requests/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NotEmpty(t, response)
	assert.False(t, response.Success)
}

func TestProtectRoleRequest(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/requests/", nil)
	req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusForbidden, w.Code)
	assert.NotEmpty(t, response)
	assert.False(t, response.Success)
}

var requestID string

func TestCreateRequest(t *testing.T) {
	w := httptest.NewRecorder()

	newRequest := request.NewRequest{
		Type:    "take_leave",
		Content: "em đưa vợ đi đẻ",
	}
	body, _ := json.Marshal(newRequest)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/requests/", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)
	data, _ := response.Data.(map[string]interface{})
	requestID = strconv.Itoa(int(data["id"].(float64)))

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}

func TestApproveRequest(t *testing.T) {
	w := httptest.NewRecorder()

	dataUpdate := map[string]interface{}{
		"status": "approve",
	}

	body, _ := json.Marshal(dataUpdate)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/requests/"+requestID, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}
