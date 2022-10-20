package department

import (
	errorModels "employee_manage/constant"
	"employee_manage/models"
	"employee_manage/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDepartmentByID godoc
// @Tags department
// @Summary Get department by ID
// @Description Get department by ID on the system
// @Param department_id path int true "id of department"
// @Success 200 {object} models.Department
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Security Authentication
// @Router /departments/{department_id} [GET]
func GetDepartmentByID(c *gin.Context) {
	var department models.Department

	departmentID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		appError := errorModels.NewAppError(errors.New("department id is invalid"), errorModels.ValidationError)
		c.Error(appError)
		return
	}

	err = models.GetDepartmentByID(&department, departmentID)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "Success",
		Data:    department,
	})
}

// CreateDepartment godoc
// @Tags department
// @Summary Create new department
// @Description Create new department on the system
// @Accept  json
// @Produce  json
// @Param data body NewDepartment true "body data"
// @Success 200 {object} models.Department
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Security Authentication
// @Router /departments [POST]
func CreateDepartment(c *gin.Context) {
	var request NewDepartment

	if err := c.BindJSON(&request); err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	errValidate := validateStructDepartment(request)

	if errValidate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   errValidate,
		})
		return
	}

	department := models.Department{
		Name:           request.Name,
		DepartmentCode: request.DepartmentCode,
		Address:        request.Address,
		Status:         request.Status,
	}

	err := models.CreateDepartment(&department)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "Success",
		Data:    department,
	})
}

// UpdateDepartmentByID godoc
// @Tags department
// @Summary Update department by ID
// @Description Update department by ID on the system
// @Param department_id path int true "id of department"
// @Param data body NewDepartment true "body data"
// @Success 200 {object} models.Department
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Security Authentication
// @Router /departments/{department_id} [PUT]
func UpdateDepartmentByID(c *gin.Context) {
	departmentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appError := errorModels.NewAppError(errors.New("department id is invalid"), errorModels.ValidationError)
		c.Error(appError)
		return
	}

	var requestMap map[string]interface{}

	err = c.ShouldBind(&requestMap)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	errValidate := updateValidation(requestMap)

	if errValidate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   errValidate,
		})
		return
	}

	department, err := models.UpdateDepartmentByID(departmentID, requestMap)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Data:    department,
	})

}

// DeleteDepartmentByID godoc
// @Tags department
// @Summary Delete department by ID
// @Description Delete department by ID on the system
// @Param department_id path int true "id of department"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Security Authentication
// @Router /departments/{department_id} [DELETE]
func DeleteDepartmentByID(c *gin.Context) {
	departmentID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		appError := errorModels.NewAppError(errors.New("user id is invalid"), errorModels.ValidationError)
		c.Error(appError)
		return
	}

	err = models.DeleteDepartmentByID(departmentID)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "Department deleted successfully",
	})
}

func ExportExcel(c *gin.Context) {
	departmentID := c.GetStringMap("role")["department_id"].(int)
	users, _ := models.GetUsersByDepartmentID(departmentID)

	if err := services.ExportExcel(strconv.Itoa(departmentID), users); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "Export success",
	})
}

func ExportExcelByDepartmentID(c *gin.Context) {
	departmentID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		appError := errorModels.NewAppError(errors.New("user id is invalid"), errorModels.ValidationError)
		c.Error(appError)
		return
	}

	users, _ := models.GetUsersByDepartmentID(departmentID)
	if err := services.ExportExcel(strconv.Itoa(departmentID), users); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "Export success",
	})
}
