package auth

import (
	"employee_manage/constant"
	"employee_manage/models"
	"employee_manage/services"
	"fmt"
	"net/http"

	errorModels "employee_manage/constant"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Tags auth
// @Summary Login
// @Description Login account
// @Accept  json
// @Produce  json
// @Param data body RequestLogin true "body data"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /auth/login [POST]
func Login(c *gin.Context) {
	var request RequestLogin

	if err := c.BindJSON(&request); err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	errValidate := validateStructLogin(request)

	if errValidate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   errValidate,
		})
		return
	}

	var user models.User
	err := models.GetUserByEmail(&user, request.Email)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	isCorrectPassword := models.CheckPasswordHash(user.Password, request.Password)
	fmt.Println(user.Password, request.Password, isCorrectPassword)

	if !isCorrectPassword {
		appError := errorModels.NewAppErrorWithType(errorModels.WrongPassword)
		c.Error(appError)
		return
	}

	payload := services.TokenPayload{
		UserID: user.ID,
	}

	accessToken, err := services.SignToken(constant.AccessToken, payload)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	refreshToken, err := services.SignToken(constant.RefreshToken, payload)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	err = models.CreateToken(&models.Token{
		Type:   constant.RefreshToken,
		UserID: &user.ID,
		Value:  refreshToken.Token,
	})

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": LoginResponse{
			AccessToken:  accessToken.Token,
			RefreshToken: refreshToken.Token,
		},
	})
}

func SendCodeResetPassword(c *gin.Context) {

	var request RequestSendCodeResetPassword
	if err := c.BindJSON(&request); err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	var user models.User
	err := models.GetUserByEmail(&user, request.Email)
	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	code, err := models.GenerateResetPassword(user, request.Email)
	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, code)
}

func ResetPassword(c *gin.Context) {

	var request RequestResetPassword
	if err := c.BindJSON(&request); err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	token, err := models.CheckToken(request.Code)
	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	err = models.ResetPassword(token, request.NewPassword)
	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	_ = models.DeleteToken(request.Code)

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "Reset password successfully",
	})
}
