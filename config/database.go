package config

import (
	"fmt"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func GormOpen() (gormDB *gorm.DB, err error) {
	config := ConfigApp.DbConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.Username, config.Password, config.Hostname, config.Port, config.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
