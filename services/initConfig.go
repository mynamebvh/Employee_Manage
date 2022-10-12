package services

import (
	"employee_manage/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitialGinConfig(router *gin.Engine) {
	var err error

	err = config.LoadEnv()

	if err != nil {
		panic(fmt.Errorf("fatal error in load env: %s", err))
	}

	config.DB, err = config.GormOpen()

	if err != nil {
		panic(fmt.Errorf("fatal error in database file: %s", err))
	}

}
