package config

import (
	"github.com/spf13/viper"
)

// DBConfig represents db configuration
type MySQLConfig struct {
	Hostname string
	Port     string
	Username string
	DBName   string
	Password string
}

type RedisConfig struct {
	Address  string
	Password string
}

type JWTConfig struct {
	SecretAccessToken   string
	ExpiredAccessToken  int
	SecretRefreshToken  string
	ExpiredRefreshToken int
}
type MailConfig struct {
	SMTP     string
	Email    string
	Port     int
	Password string
}
type ConfigApplication struct {
	Environment string
	ServerPort  string
	MySQL       MySQLConfig
	JWT         JWTConfig
	Redis       RedisConfig
	Mail        MailConfig
}

var ConfigApp ConfigApplication

func LoadEnv() (err error) {
	viper.SetConfigFile("config.json")
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	ConfigApp.ServerPort = viper.GetString("ServerPort")

	err = viper.Unmarshal(&ConfigApp)

	return
}
