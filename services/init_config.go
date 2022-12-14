package services

import (
	"employee_manage/config"
	"fmt"
)

func InitialGinConfig() {
	var err error

	err = config.LoadEnv()

	if err != nil {
		panic(fmt.Errorf("fatal error in load env: %s", err))
	}

	config.DB, err = config.GormOpen()

	if err != nil {
		panic(fmt.Errorf("fatal error in database file: %s", err))
	}

	config.Rdb, err = config.ConnectRedis()

	if err != nil {
		panic(fmt.Errorf("error connect redis: %s", err))
	}
}
