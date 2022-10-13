package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"employee_manage/config"
	middlewares "employee_manage/middlewares"
	"employee_manage/routes"
	"employee_manage/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	services.InitialGinConfig(router)
	router.Use(middlewares.HandlerError)
	routes.ApplicationV1Router(router)
	startServer(router)
}

func startServer(router http.Handler) {
	serverPort := fmt.Sprintf(":%s", config.ConfigApp.ServerPort)
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
