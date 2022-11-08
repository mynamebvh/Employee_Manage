package config

import (
	"bytes"
	"employee_manage/controllers/auth"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

var JWTAccountRoot auth.LoginResponse
var JWTAccountManager auth.LoginResponse
var JWTAccountUser auth.LoginResponse

type RequestLogin struct {
	ID       int
	Email    string
	Password string
}

func GetAccountRoot() RequestLogin {
	return RequestLogin{
		ID:       1,
		Email:    "mynamebvh@gmail.com",
		Password: "hoangdz",
	}
}

func GetAccountManager() RequestLogin {
	return RequestLogin{
		ID:       2,
		Email:    "mynamebvh1@gmail.com",
		Password: "hoangdz",
	}
}

func GetAccountUser() RequestLogin {
	return RequestLogin{
		ID:       3,
		Email:    "mynamebvh2@gmail.com",
		Password: "hoangdz",
	}
}

type ResponseLogin struct {
	Success bool
	Data    auth.LoginResponse
}

func CallAPILogin(account RequestLogin, r *gin.Engine) auth.LoginResponse {
	rows := Mock.NewRows([]string{"id", "email", "password"}).
		AddRow(account.ID, account.Email, "$2a$12$oGIJ9QJ43E6INO4usXcoGebRB7N1lfXKOVPpnj6.vzQcPnOm9SfDe")
	Mock.ExpectQuery("SELECT (.+) FROM `users` WHERE email=?(.+)").WithArgs(account.Email).WillReturnRows(rows)
	Mock.ExpectBegin()
	Mock.ExpectExec("INSERT INTO `tokens` (.+)").WillReturnResult(sqlmock.NewResult(1, 1))
	Mock.ExpectCommit()

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
