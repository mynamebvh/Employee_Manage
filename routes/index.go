package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "employee_manage/docs"
)

// @title Employee Manager
// @version 1.0
// @description Documentation's Employee Manager
// @termsOfService http://swagger.io/terms/

// @contact.name Bui Viet Hoang
// @contact.url
// @contact.email mynamebvh@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v1
func ApplicationV1Router(router *gin.Engine) {

	router.GET("/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")
	{

		// User route
		v1HelloWorld := v1.Group("/users")
		userRoute(v1HelloWorld)
	}
}
