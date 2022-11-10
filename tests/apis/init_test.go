package apis

import (
	"employee_manage/middlewares"
	. "employee_manage/tests/config"

	"employee_manage/routes"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

type TestModel struct {
	Name  string
	Type  string
	Args  interface{}
	Mock  func()
	Want  interface{}
	Error bool
}

func setupTest() {
	router = gin.Default()
	Mock = InitialTestConfig()
	router.Use(middlewares.HandlerError)
	routes.ApplicationV1Router(router)
}

func TestMain(m *testing.M) {
	setupTest()
	SetJwt()
	os.Exit(m.Run())
}
