package routes

import (
	controllers "employee_manage/controllers/department"
	"employee_manage/middlewares"

	"github.com/gin-gonic/gin"
)

func departmentRoute(route *gin.RouterGroup) {
	route.GET("/:id", middlewares.ProtectRole([]string{"admin"}, "requests"), controllers.GetDepartmentByID)
	route.POST("/", controllers.CreateDepartment)
	route.POST("/export-excel/:id", controllers.ExportExcel)
	route.PUT("/:id", controllers.UpdateDepartmentByID)
	route.DELETE("/:id", controllers.DeleteDepartmentByID)
}
