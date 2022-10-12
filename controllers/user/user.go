package user

import (
	"errors"
	"net/http"
	"strconv"

	errorModels "employee_manage/constant"
	"employee_manage/models"

	"github.com/gin-gonic/gin"
)

// GetUserByID godoc
// @Tags user
// @Summary Get user by ID
// @Description Get user by ID on the system
// @Param user_id path int true "id of user"
// @Success 200 {object} models.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users/{user_id} [GET]
func GetUserByID(c *gin.Context) {
	var user models.User

	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		appError := errorModels.NewAppError(errors.New("user id is invalid"), errorModels.ValidationError)
		c.Error(appError)
		return
	}

	err = models.GetUserByID(&user, userID)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Tags user
// @Summary Create New User
// @Description Create new user on the system
// @Accept  json
// @Produce  json
// @Param data body NewUser true "body data"
// @Success 200 {object} models.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users [POST]
func CreateUser(c *gin.Context) {
	var request NewUser

	if err := c.BindJSON(&request); err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	errValidate := validateStructUser(request)

	if errValidate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   errValidate,
		})
		return
	}

	user := models.User{
		EmployeeCode: request.EmployeeCode,
		FullName:     request.FullName,
		Password:     request.Password,
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
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserByID godoc
// @Tags user
// @Summary Update user by ID
// @Description Update user by ID on the system
// @Param user_id path int true "id of user"
// @Param data body NewUser true "body data"
// @Success 200 {object} models.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users/{user_id} [PUT]
func UpdateUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appError := errorModels.NewAppError(errors.New("user id is invalid"), errorModels.ValidationError)
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

	// errValidate := updateValidation(requestMap)

	// if errValidate != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"success": false,
	// 		"error":   errValidate,
	// 	})
	// 	return
	// }

	user, err := models.UpdateUserByID(userID, requestMap)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)

}

// DeleteUserByID godoc
// @Tags user
// @Summary Delete user by ID
// @Description Delete user by ID on the system
// @Param user_id path int true "id of user"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /users/{user_id} [DELETE]
func DeleteUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		appError := errorModels.NewAppError(errors.New("user id is invalid"), errorModels.ValidationError)
		c.Error(appError)
		return
	}

	err = models.DeleteUserByID(userID)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}

func ChangePassword(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		appError := errorModels.NewAppError(errors.New("user id is invalid"), errorModels.ValidationError)
		c.Error(appError)
		return
	}

	var user models.User
	err = models.GetUserByID(&user, userID)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	var request RequestChangePassword
	if err := c.BindJSON(&request); err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	errValidate := validateStructChangePassword(request)
	if errValidate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   errValidate,
		})
		return
	}

	err = models.ChangePassword(user, request.OldPassword, request.NewPassword)
	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "Change password successfully",
	})
}
