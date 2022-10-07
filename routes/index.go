package routes

import (
	"github.com/gin-gonic/gin"
)

func ApplicationV1Router(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// User route
		v1HelloWorld := v1.Group("/users")
		userRoute(v1HelloWorld)
	}
}
