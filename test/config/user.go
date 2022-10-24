package config

import (
	"bytes"
	"employee_manage/controllers/auth"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var JWTAccountRoot auth.LoginResponse
var JWTAccountManager auth.LoginResponse
var JWTAccountUser auth.LoginResponse

func GetAccountRoot() auth.RequestLogin {
	return auth.RequestLogin{
		Email:    "mynamebvh@gmail.com",
		Password: "hoangdz1",
	}
}

func GetAccountManager() auth.RequestLogin {
	return auth.RequestLogin{
		Email:    "mynamebvh1@gmail.com",
		Password: "hoangdz",
	}
}

func GetAccountUser() auth.RequestLogin {
	return auth.RequestLogin{
		Email:    "mynamebvh2@gmail.com",
		Password: "hoangdz",
	}
}

type ResponseLogin struct {
	Success bool
	Data    auth.LoginResponse
}

func CallAPILogin(account auth.RequestLogin, r *gin.Engine) auth.LoginResponse {
	w := httptest.NewRecorder()
	body, _ := json.Marshal(account)

	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)
	var response ResponseLogin
	json.Unmarshal(w.Body.Bytes(), &response)
	return response.Data
}

func SetJwt() {

	router := gin.Default()
	router.POST("/login", auth.Login)

	JWTAccountRoot = CallAPILogin(GetAccountRoot(), router)
	JWTAccountManager = CallAPILogin(GetAccountManager(), router)
	JWTAccountUser = CallAPILogin(GetAccountUser(), router)
}
