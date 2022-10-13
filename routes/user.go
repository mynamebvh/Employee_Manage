package routes

import (
	controllers "employee_manage/controllers/user"
	"employee_manage/middlewares"

	"github.com/gin-gonic/gin"
)

func userRoute(route *gin.RouterGroup) {
	route.GET("/me", controllers.GetMe)
	route.GET("/:id", middlewares.ProtectRole([]string{"user"}), controllers.GetUserByID)
	route.POST("/", controllers.CreateUser)
	route.PUT("/change-password/:id", controllers.ChangePassword)
	route.PUT("/:id", controllers.UpdateUserByID)
	route.DELETE("/:id", controllers.DeleteUserByID)
}
