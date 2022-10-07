package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Hostname string
	Port     string
	Username string
	DBName   string
	Password string
}

func GormOpen() (gormDB *gorm.DB, err error) {
	viper.SetConfigFile("config.json")
	err = viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	configDB := viper.GetStringMap("Databases.MYSQL")

	var config DBConfig
	errDecode := mapstructure.Decode(configDB, &config)

	if errDecode != nil {
		return nil, errDecode
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.Username, config.Password, config.Hostname, config.Port, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
