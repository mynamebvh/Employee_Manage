package controllers

import (
	"employee_manage/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserByID(c *gin.Context) {
	var user models.User

	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return
	}

	err = models.GetUserByID(&user, userID)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var request NewUser

	if err := c.BindJSON(&request); err != nil {
		return
	}

	user := models.User{
		EmployeeCode: request.EmployeeCode,
		FullName:     request.FullName,
		Phone:        request.Phone,
		Email:        request.Email,
		Gender:       request.Gender,
		Birthday:     request.Birthday,
		Address:      request.Address,
		DepartmentID: request.DepartmentID,
		RoleID:       request.RoleID,
	}

	err := models.CreateUser(&user)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	var requestMap map[string]interface{}

	err = c.ShouldBind(&requestMap)

	user, err := models.UpdateUserByID(userID, requestMap)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, user)

}

func DeleteUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		return
	}

	err = models.DeleteUserByID(userID)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
