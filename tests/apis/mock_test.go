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

// func TestRootGetUserByID(t *testing.T) {
// 	mocks.MockGetUserByID()
// 	w := httptest.NewRecorder()

// 	req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/3", nil)
// 	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
// 	router.ServeHTTP(w, req)

// 	var response user.MessageResponse
// 	json.Unmarshal(w.Body.Bytes(), &response)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.NotEmpty(t, response)
// 	assert.True(t, response.Success)
// }

// func TestRootGetUsers(t *testing.T) {
// 	tests := []TestModel{
// 		{
// 			Name: "Test Root Get Users",
// 			Type: "Success",
// 			Args: nil,
// 			API: API{
// 				URL:    "/api/v1/users/",
// 				Method: http.MethodGet,
// 			},
// 			Mock: mocks.MockGetUsers,
// 			Want: http.StatusOK,
// 		},
// 		{
// 			Name: "Test Root Get User By ID",
// 			Type: "Success",
// 			Args: nil,
// 			API: API{
// 				URL:    "/api/v1/users/3",
// 				Method: http.MethodGet,
// 			},
// 			Mock: mocks.MockGetUserByID,
// 			Want: http.StatusOK,
// 		},
// 		{
// 			Name: "Test Protect Role Get Users",
// 			Type: "Role",
// 			Args: nil,
// 			API: API{
// 				URL:    "/api/v1/users/",
// 				Method: http.MethodPost,
// 			},
// 			Mock: mocks.MockUser,
// 			Want: http.StatusForbidden,
// 		},
// 	}

// 	for _, tt := range tests {
// 		w := httptest.NewRecorder()
// 		tt.Mock()
// 		body, _ := json.Marshal(tt.Args)
// 		req, _ := http.NewRequest(tt.API.Method, tt.API.URL, bytes.NewBuffer(body))
// 		if tt.Type == "Role" {
// 			req.Header.Set("Authorization", "Bearer "+JWTAccountUser.AccessToken)
// 		} else {
// 			req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
// 		}
// 		req.Header.Set("Content-Type", "application/json")
// 		router.ServeHTTP(w, req)

// 		var response auth.MessageResponse
// 		json.Unmarshal(w.Body.Bytes(), &response)

// 		if tt.Type == "Success" {
// 			assert.Equal(t, tt.Want, w.Code)
// 			assert.NotEmpty(t, response)
// 			assert.True(t, response.Success)
// 		} else if tt.Type == "Fail" {
// 			assert.Equal(t, tt.Want, w.Code)
// 			assert.NotEmpty(t, response)
// 			assert.False(t, response.Success)
// 		} else {
// 			var response map[string]interface{}
// 			json.Unmarshal(w.Body.Bytes(), &response)
// 			assert.Equal(t, tt.Want, w.Code)
// 			assert.NotEmpty(t, response)
// 			assert.False(t, response["success"].(bool))
// 		}
// 	}
// }

// func TestRootCreateUser(t *testing.T) {
// 	mocks.MockCreateUser()
// 	w := httptest.NewRecorder()
// 	birthday, _ := time.Parse("2006-01-02 15:04:05.000 -0700", "2020-01-02 03:04:05.000 +0000")

// 	newUser := user.NewUser{
// 		FullName:     "Bui Viet Hoang",
// 		Password:     "hoangdz",
// 		Phone:        "0966150922",
// 		Email:        "mynamebvh19@gmail.com",
// 		Gender:       true,
// 		Address:      "Hà Nội",
// 		DepartmentID: 1,
// 		RoleID:       3,
// 		Birthday:     birthday,
// 	}
// 	body, _ := json.Marshal(newUser)

// 	req, _ := http.NewRequest(http.MethodPost, "/api/v1/users/", bytes.NewBuffer(body))
// 	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
// 	req.Header.Set("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)

// 	var response user.MessageResponse
// 	json.Unmarshal(w.Body.Bytes(), &response)

// 	assert.Equal(t, http.StatusCreated, w.Code)
// 	assert.NotEmpty(t, response)
// 	assert.True(t, response.Success)
// }

// func TestRootDeleteUserByID(t *testing.T) {
// 	mocks.MockDeleteUserByID()
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/users/4", nil)
// 	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
// 	router.ServeHTTP(w, req)

// 	var response user.MessageResponse
// 	json.Unmarshal(w.Body.Bytes(), &response)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.NotEmpty(t, response)
// 	assert.True(t, response.Success)
// }

func TestRootUpdateUserByID(t *testing.T) {
	mocks.MockUpdateUserByID()
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

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/users/4", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+JWTAccountRoot.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var response user.MessageResponse
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, response)
	assert.True(t, response.Success)
}

// func TestUserSendMailForgetPassword(t *testing.T) {
// 	mocks.MockSendCode()
// 	w := httptest.NewRecorder()

// 	data := map[string]interface{}{
// 		"email": "mynamebvh4@gmail.com",
// 	}

// 	body, _ := json.Marshal(data)
// 	req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/send-code", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)

// 	fmt.Println(w.Body.String())

// 	var response user.MessageResponse
// 	json.Unmarshal(w.Body.Bytes(), &response)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.NotEmpty(t, response)
// 	assert.True(t, response.Success)
// }
