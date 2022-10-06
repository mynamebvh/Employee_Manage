package routes

import (
	"github.com/gin-gonic/gin"
)

func ApplicationV1Router(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// Hello World
		v1HelloWorld := v1.Group("/hello")
		helloRoute(v1HelloWorld)
	}
}
