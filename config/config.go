package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// DBConfig represents db configuration
type DBConfig struct {
	Hostname string
	Port     string
	Username string
	DBName   string
	Password string
}

type JWTConfig struct {
	SecretAccessToken   string
	ExpiredAccessToken  int
	SecretRefreshToken  string
	ExpiredRefreshToken int
}

type ConfigApplication struct {
	ServerPort string
	DbConfig   DBConfig
	JwtConfig  JWTConfig
}

var ConfigApp ConfigApplication

func LoadEnv() (err error) {
	viper.SetConfigFile("config.json")
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	ConfigApp.ServerPort = viper.GetString("ServerPort")

	envDB := viper.GetStringMap("Databases.MYSQL")
	envJWT := viper.GetStringMap("JWT")

	var dbConfig DBConfig
	var jwtConfig JWTConfig

	err = mapstructure.Decode(envDB, &dbConfig)

	if err != nil {
		return
	}

	err = mapstructure.Decode(envJWT, &jwtConfig)

	if err != nil {
		return
	}

	ConfigApp.DbConfig = dbConfig
	ConfigApp.JwtConfig = jwtConfig
	return
}
