package apis

import (
	"employee_manage/controllers/user"
	. "employee_manage/tests/config"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIn(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/calendars/checkin", nil)
	req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	fmt.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}

func TestCheckOut(t *testing.T) {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/calendars/checkout", nil)
	req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	fmt.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}
