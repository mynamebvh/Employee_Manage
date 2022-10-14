package routes

import (
	controllers "employee_manage/controllers/auth"

	"github.com/gin-gonic/gin"
)

func authRoute(route *gin.RouterGroup) {
	route.POST("/login", controllers.Login)
	route.POST("/refresh-token", controllers.RefreshToken)
	route.POST("/send-code", controllers.SendCodeResetPassword)
	route.POST("/reset-password", controllers.ResetPassword)
}
