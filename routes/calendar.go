package routes

import (
	controllers "employee_manage/controllers/calendar"

	"github.com/gin-gonic/gin"
)

func calendarRoute(route *gin.RouterGroup) {
	route.POST("/checkin", controllers.CheckIn)
	route.POST("/checkout", controllers.Checkout)
	route.POST("/working-time", controllers.GetWorkingTimeInMonth)
}
