package auth

import (
	"employee_manage/config"
	"employee_manage/constant"
	"employee_manage/models"
	"employee_manage/services"
	"errors"
	"fmt"
	"net/http"

	errorModels "employee_manage/constant"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

	user, err := models.GetUserByEmail(request.Email)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	isCorrectPassword := models.CheckPasswordHash(user.Password, request.Password)

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

// Send code to mail godoc
// @Tags auth
// @Summary Send code to mail
// @Description Send code to mail
// @Accept  json
// @Produce  json
// @Param data body RequestSendCodeResetPassword true "body data"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /auth/send-code [POST]
func SendCodeResetPassword(c *gin.Context) {

	var request RequestSendCodeResetPassword
	if err := c.BindJSON(&request); err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	user, err := models.GetUserByEmail(request.Email)
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

	services.SendMail(user.Email, "Forget passowrd", fmt.Sprintf("Mã bí mật của bạn là: %s", code))
	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "submitted successfully",
	})
}

// ResetPassword godoc
// @Tags auth
// @Summary Reset password
// @Description Reset pasword
// @Accept  json
// @Produce  json
// @Param data body RequestResetPassword true "body data"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /auth/reset-password [POST]
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

	err = models.ResetPassword(token.Email, request.NewPassword)
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

func RefreshToken(c *gin.Context) {
	var request RequestRefreshToken

	if err := c.BindJSON(&request); err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	accessToken, err := jwt.Parse(request.AccessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ConfigApp.JWT.SecretAccessToken), nil
	})

	if errors.Is(err, jwt.ErrTokenMalformed) || errors.Is(err, jwt.ErrSignatureInvalid) {
		appError := errorModels.NewAppError(errors.New("access token invalid"), errorModels.NotFound)
		_ = c.Error(appError)
		return
	}

	_, err = jwt.Parse(request.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ConfigApp.JWT.SecretRefreshToken), nil
	})

	if err != nil {
		appError := errorModels.NewAppError(errors.New("refresh token invalid"), errorModels.NotFound)
		_ = c.Error(appError)
		return
	}

	_, errQuery := models.GetTokenByValueAndType(request.RefreshToken, errorModels.RefreshToken)

	if errQuery != nil {
		appError := errorModels.NewAppError(errQuery, errorModels.NotFound)
		_ = c.Error(appError)
		return
	}

	claims := accessToken.Claims.(jwt.MapClaims)
	userID := int(claims["payload"].(map[string]interface{})["user_id"].(float64))

	payload := services.TokenPayload{
		UserID: userID,
	}

	newAccessToken, _ := services.SignToken(constant.AccessToken, payload)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Data:    newAccessToken,
	})
}
