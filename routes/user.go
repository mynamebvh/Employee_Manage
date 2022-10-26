package routes

import (
	controllers "employee_manage/controllers/user"
	"employee_manage/middlewares"

	"github.com/gin-gonic/gin"
)

func userRoute(route *gin.RouterGroup) {
	route.GET("/me", controllers.GetMe)
	route.GET("/", middlewares.ProtectRole([]string{"admin"}, ""), controllers.GetUsers)
	route.GET("/:id", middlewares.ProtectRole([]string{"admin", "manager"}, "users"), controllers.GetUserByID)
	route.POST("/reset-password", controllers.ResetPassWordQueue)
	route.POST("/", middlewares.ProtectRole([]string{"admin"}, ""), controllers.CreateUser)
	route.PUT("/change-password/:id", controllers.ChangePassword)
	route.PUT("/:id", middlewares.ProtectRole([]string{"admin"}, ""), controllers.UpdateUserByID)
	route.DELETE("/:id", middlewares.ProtectRole([]string{"admin"}, ""), controllers.DeleteUserByID)
}
