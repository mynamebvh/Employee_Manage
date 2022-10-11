package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"employee_manage/config"
	middlewares "employee_manage/middlewares"
	"employee_manage/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	router := gin.Default()
	initialGinConfig(router)
	router.Use(middlewares.HandlerError)
	routes.ApplicationV1Router(router)
	startServer(router)

}

func initialGinConfig(router *gin.Engine) {
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

func startServer(router http.Handler) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error in config file: %s", err))
	}
	serverPort := fmt.Sprintf(":%s", viper.GetString("ServerPort"))
	s := &http.Server{
		Addr:           serverPort,
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		panic(fmt.Errorf("fatal error description: %s", strings.ToLower(err.Error())))
	}
}
