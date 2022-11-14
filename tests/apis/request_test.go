package apis

import (
	"bytes"
	"employee_manage/controllers/department"
	"employee_manage/controllers/request"
	. "employee_manage/tests/config"
	"employee_manage/tests/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequests(t *testing.T) {

	tests := []TestModel{
		{
			Name: "Test Get Requests",
			Type: "Success",
			Args: nil,
			Mock: mocks.MockGetRequests,
			Want: http.StatusOK,
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
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/requests/", bytes.NewBuffer(body))
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

func TestGetRequestByID(t *testing.T) {
	tests := []TestModel{
		{
			Name: "Test Get Request By ID",
			Type: "Success",
			Args: nil,
			Mock: mocks.MockGetRequestByID,
			Want: http.StatusOK,
		},
		{
			Name: "Test Fail Get Request By ID",
			Type: "Fail",
			Args: nil,
			Mock: mocks.MockFailGetRequestByID,
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
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/requests/1", bytes.NewBuffer(body))
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
			t.Log("err", response)
			assert.Equal(t, tt.Want, w.Code)
			assert.NotEmpty(t, response)
			assert.False(t, response.Success)
		}
	}
}

func TestCreateRequest(t *testing.T) {
	tests := []TestModel{
		{
			Name: "Test Root Create Request",
			Type: "Success",
			Args: request.NewRequest{
				Type:    "take_leave",
				Content: "em đưa vợ đi đẻ",
			},
			Mock: mocks.MockCreateRequest,
			Want: http.StatusCreated,
		},
		{
			Name: "Test Auth",
			Type: "Auth",
			Args: nil,
			Mock: func() {

			},
			Want: http.StatusUnauthorized,
		},
		{
			Name: "Test Validate",
			Type: "Validate",
			Args: request.NewRequest{
				Type: "take_leave",
			},
			Mock: func() {},
			Want: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		tt.Mock()
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tt.Args)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/requests/", bytes.NewBuffer(body))
		if tt.Type != "Auth" {
			req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
		}
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		var response department.MessageResponse
		json.Unmarshal(w.Body.Bytes(), &response)

		t.Log("err", response)
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

func TestApproveRequest(t *testing.T) {
	tests := []TestModel{
		{
			Name: "Test Approve Request By ID",
			Type: "Success",
			Args: map[string]interface{}{
				"status": "approve",
			},
			Mock: mocks.MockUpdateRequestByID,
			Want: http.StatusOK,
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
		req, _ := http.NewRequest(http.MethodPut, "/api/v1/requests/1", bytes.NewBuffer(body))
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
