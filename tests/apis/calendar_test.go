package apis

import (
	"bytes"
	"employee_manage/controllers/user"
	. "employee_manage/tests/config"
	"employee_manage/tests/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIn(t *testing.T) {
	tests := []TestModel{
		{
			Name: "Test CheckIn Success",
			Type: "Success",
			Args: nil,
			Mock: mocks.MockCheckIn,
			Want: http.StatusOK,
		},
		{
			Name: "Test CheckIn Fail",
			Type: "Fail",
			Args: nil,
			Mock: func() {},
			Want: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		tt.Mock()
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tt.Args)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/calendars/checkin", bytes.NewBuffer(body))
		if tt.Type == "Success" {
			req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
		}

		router.ServeHTTP(w, req)

		var response user.MessageResponse
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

func TestCheckOut(t *testing.T) {

	tests := []TestModel{
		{
			Name: "Test CheckOut Success",
			Type: "Success",
			Args: nil,
			Mock: mocks.MockCheckOut,
			Want: http.StatusOK,
		},
		{
			Name: "Test CheckOut Fail",
			Type: "Fail",
			Args: nil,
			Mock: func() {},
			Want: http.StatusUnauthorized,
		},
	}

	for _, tt := range tests {
		tt.Mock()
		w := httptest.NewRecorder()
		body, _ := json.Marshal(tt.Args)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/calendars/checkout", bytes.NewBuffer(body))
		if tt.Type == "Success" {
			req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
		}

		router.ServeHTTP(w, req)

		var response user.MessageResponse
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
