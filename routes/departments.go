package routes

import (
	controllers "employee_manage/controllers/department"
	"employee_manage/middlewares"

	"github.com/gin-gonic/gin"
)

func departmentRoute(route *gin.RouterGroup) {
	route.GET("/", middlewares.ProtectRole([]string{"admin"}, ""), controllers.GetDepartments)
	route.GET("/:id", middlewares.ProtectRole([]string{"admin"}, "requests"), controllers.GetDepartmentByID)
	route.POST("/", middlewares.ProtectRole([]string{"admin"}, ""), controllers.CreateDepartment)
	route.POST("/export-excel", middlewares.ProtectRole([]string{"manager"}, ""), controllers.ExportExcel)
	route.POST("/export-excel/:id", middlewares.ProtectRole([]string{"admin"}, ""), controllers.ExportExcelByDepartmentID)
	route.PUT("/:id", middlewares.ProtectRole([]string{"admin"}, ""), controllers.UpdateDepartmentByID)
	route.DELETE("/:id", middlewares.ProtectRole([]string{"admin"}, ""), controllers.DeleteDepartmentByID)
}
