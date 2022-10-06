package routes

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloRoute(route *gin.RouterGroup) {
	route.GET("/", func(ctx *gin.Context) {
		errors.New("Something unexpected happend!")
		ctx.String(http.StatusOK, "hello")
	})
}
