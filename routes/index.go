package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "employee_manage/docs"
	"employee_manage/middlewares"
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
// @BasePath /api/v1
func ApplicationV1Router(router *gin.Engine) {

	router.GET("/v1/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api/v1")
	{
		// Auth
		v1Auth := v1.Group("/auth")
		authRoute(v1Auth)

		// User route
		v1User := v1.Group("/users", middlewares.Protect())
		userRoute(v1User)

		// Department route
		v1Department := v1.Group("/departments", middlewares.Protect())
		departmentRoute(v1Department)
	}
}
