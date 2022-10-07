package routes

import (
	controllers "employee_manage/controllers/user"

	"github.com/gin-gonic/gin"
)

func userRoute(route *gin.RouterGroup) {
	route.GET("/:id", controllers.GetUserByID)
	route.POST("/", controllers.CreateUser)
	route.PUT("/:id", controllers.UpdateUserByID)
	route.DELETE("/:id", controllers.DeleteUserByID)
}
