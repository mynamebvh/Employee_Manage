package user

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"employee_manage/config"
	errorModels "employee_manage/constant"
	"employee_manage/models"
	"employee_manage/utils"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	data, err := models.GetUsers(c)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"page_current": data.Current,
		"page_size":    data.PageSize,
		"total":        data.Total,
		"data":         data.Data,
	})

}

// GetMe godoc
// @Tags user
// @Summary Get info user current
// @Description Get info user current on the system
// @Success 200 {object} dto.QueryResultGetMe
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Security Authentication
// @Router /users/me [GET]
func GetMe(c *gin.Context) {
	var user models.User

	userID := utils.GetUserIDByContext(c)

	data, err := models.GetMe(&user, userID)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, data)
}

// GetUserByID godoc
// @Tags user
// @Summary Get user by ID
// @Description Get user by ID on the system
// @Param user_id path int true "id of user"
// @Success 200 {object} models.User
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Security Authentication
// @Router /users/{user_id} [GET]
func GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		appError := errorModels.NewAppError(errors.New("user id is invalid"), errorModels.ValidationError)
		c.Error(appError)
		return
	}

	user, err := models.GetUserByID(userID)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Data:    user,
	})
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
// @Security Authentication
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
		EmployeeCode: fmt.Sprintf("NV%d", rand.Intn(5000-1)+1),
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

	c.JSON(http.StatusCreated, MessageResponse{
		Success: true,
		Data:    user,
	})
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
// @Security Authentication
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

	errValidate := updateValidation(requestMap)

	if errValidate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   errValidate,
		})
		return
	}

	user, err := models.UpdateUserByID(userID, requestMap)

	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "Update success",
		Data:    user,
	})

}

// DeleteUserByID godoc
// @Tags user
// @Summary Delete user by ID
// @Description Delete user by ID on the system
// @Param user_id path int true "id of user"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Security Authentication
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

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "user deleted successfully",
	})
}

// ChangePassword godoc
// @Tags user
// @Summary Change password
// @Description Change password on the system
// @Param data body RequestChangePassword true "body data"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Security Authentication
// @Router /users/change-password [PUT]
func ChangePassword(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		appError := errorModels.NewAppError(errors.New("user id is invalid"), errorModels.ValidationError)
		c.Error(appError)
		return
	}

	user, err := models.GetUserByID(userID)

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

func ResetPassWordQueue(c *gin.Context) {
	var request []string

	if err := c.BindJSON(&request); err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	err := config.Publish("reset-email", request)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "Successfully",
	})
}
