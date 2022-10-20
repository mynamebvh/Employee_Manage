package routes

import (
	controllers "employee_manage/controllers/request"
	"employee_manage/middlewares"

	"github.com/gin-gonic/gin"
)

func requestRoute(route *gin.RouterGroup) {
	route.GET("/", middlewares.ProtectRole([]string{"admin", "manager"}, ""), controllers.GetRequests)
	route.GET("/:id", middlewares.ProtectRole([]string{"admin", "manager"}, "requests"), controllers.GetRequestByID)
	route.POST("/", controllers.CreateRequest)
	route.PUT("/:id", middlewares.ProtectRole([]string{"admin", "manager"}, "requests"), controllers.UpdateRequestByID)
}
