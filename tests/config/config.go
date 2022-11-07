package config

import (
	"employee_manage/config"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mock sqlmock.Sqlmock

func LoadEnvTest() (err error) {
	viper.SetConfigFile("../../config.json")
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	config.ConfigApp.ServerPort = viper.GetString("ServerPort")

	err = viper.Unmarshal(&config.ConfigApp)

	return
}

func InitialTestConfig() sqlmock.Sqlmock {
	var err error

	err = LoadEnvTest()

	if err != nil {
		panic(fmt.Errorf("fatal error in load env: %s", err))
	}

	db, mock, err := sqlmock.New()

	if err != nil {
		fmt.Printf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
		// DriverName: "mysql",
	})

	config.DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		fmt.Printf("Error gorm connect mock %s", err)
	}
	return mock
	// if err != nil {
	// 	panic(fmt.Errorf("fatal error in database file: %s", err))
	// }

	// config.Rdb, err = config.ConnectRedis()

	// if err != nil {
	// 	panic(fmt.Errorf("error connect redis: %s", err))
	// }
}
