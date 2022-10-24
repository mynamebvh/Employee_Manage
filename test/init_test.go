package test

import (
	"employee_manage/middlewares"
	. "employee_manage/test/config"

	"employee_manage/routes"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func setupTest() {
	router = gin.Default()
	InitialGinConfig()
	router.Use(middlewares.HandlerError)
	routes.ApplicationV1Router(router)
}

func TestMain(m *testing.M) {
	setupTest()
	SetJwt()
	os.Exit(m.Run())
}
