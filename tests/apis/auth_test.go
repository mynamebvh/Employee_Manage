package apis

import (
	"bytes"
	"employee_manage/controllers/auth"
	"employee_manage/tests/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// import (
// 	"bytes"
// 	"employee_manage/controllers/auth"
// 	"employee_manage/tests/mocks"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

func TestLoginRoute(t *testing.T) {
	tests := []TestModel{
		{
			Name: "Test login success",
			Type: "Success",
			Args: auth.RequestLogin{
				Email:    "mynamebvh@gmail.com",
				Password: "hoangdz",
			},
			Mock: mocks.MockAuth,
			Want: http.StatusOK,
		},
		{
			Name: "Test login fail",
			Type: "Fail",
			Args: auth.RequestLogin{
				Email:    "mynamebvh@gmail.com",
				Password: "hoangdz1",
			},
			Mock: mocks.MockAuthFail,
			Want: http.StatusUnauthorized,
		},
		{
			Name: "Test validate login",
			Type: "Validate",
			Args: auth.RequestLogin{
				Email: "mynamebvh@gmail.com",
			},
			Mock: func() {

			},
			Want: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		w := httptest.NewRecorder()
		tt.Mock()
		body, _ := json.Marshal(tt.Args)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		var response auth.MessageResponse
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
			assert.Equal(t, http.StatusBadRequest, w.Code)
			assert.NotEmpty(t, response)
			assert.False(t, response["success"].(bool))
		}
	}
}

// func TestErrorParseJson(t *testing.T) {
// 	w := httptest.NewRecorder()

// 	login := map[string]interface{}{
// 		"email":    1,
// 		"password": "hoangdz",
// 	}
// 	body, _ := json.Marshal(login)

// 	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusBadRequest, w.Code)
// 	assert.Empty(t, w.Body.Bytes())
// }
