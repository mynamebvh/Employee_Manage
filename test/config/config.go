package config

import (
	"employee_manage/config"
	"fmt"

	"github.com/spf13/viper"
)

func LoadEnvTest() (err error) {
	viper.SetConfigFile("../config.json")
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	config.ConfigApp.ServerPort = viper.GetString("ServerPort")

	err = viper.Unmarshal(&config.ConfigApp)

	return
}

func InitialGinConfig() {
	var err error

	err = LoadEnvTest()

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
