package routes

import (
	controllers "employee_manage/controllers/calendar"
	"employee_manage/middlewares"

	"github.com/gin-gonic/gin"
)

func calendarRoute(route *gin.RouterGroup) {
	route.POST("/checkin", middlewares.Protect(), controllers.CheckIn)
	route.POST("/checkout", middlewares.Protect(), controllers.Checkout)
	route.POST("/working-time", middlewares.Protect(), controllers.GetWorkingTimeInMonth)
}
