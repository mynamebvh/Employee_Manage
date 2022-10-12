package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	middlewares "employee_manage/middlewares"
	"employee_manage/routes"
	"employee_manage/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	router := gin.Default()
	services.InitialGinConfig(router)
	router.Use(middlewares.HandlerError)
	routes.ApplicationV1Router(router)
	startServer(router)

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
