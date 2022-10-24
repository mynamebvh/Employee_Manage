package test

import (
	"bytes"
	"employee_manage/controllers/auth"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginRoute(t *testing.T) {
	w := httptest.NewRecorder()

	login := auth.RequestLogin{
		Email:    "mynamebvh@gmail.com",
		Password: "hoangdz1",
	}

	body, _ := json.Marshal(login)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	var response auth.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}

func TestWrongPassword(t *testing.T) {
	w := httptest.NewRecorder()

	login := auth.RequestLogin{
		Email:    "mynamebvh@gmail.com",
		Password: "hoangdz",
	}
	body, _ := json.Marshal(login)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	var response auth.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NotEmpty(t, response)
	assert.False(t, response.Success)
}

func TestValidateLogin(t *testing.T) {
	w := httptest.NewRecorder()

	login := auth.RequestLogin{
		Password: "hoangdz",
	}
	body, _ := json.Marshal(login)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, response)
	assert.False(t, response["success"].(bool))
}

func TestErrorParseJson(t *testing.T) {
	w := httptest.NewRecorder()

	login := map[string]interface{}{
		"email":    1,
		"password": "hoangdz",
	}
	body, _ := json.Marshal(login)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Empty(t, w.Body.Bytes())
}
