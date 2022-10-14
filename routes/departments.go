package routes

import (
	controllers "employee_manage/controllers/department"
	"employee_manage/middlewares"

	"github.com/gin-gonic/gin"
)

func departmentRoute(route *gin.RouterGroup) {
	route.GET("/:id", middlewares.ProtectRole([]string{"admin"}), controllers.GetDepartmentByID)
	route.POST("/", controllers.CreateDepartment)
	// route.PUT("/change-password/:id", controllers.ChangePassword)
	route.PUT("/:id", controllers.UpdateDepartmentByID)
	route.DELETE("/:id", controllers.DeleteDepartmentByID)
}
