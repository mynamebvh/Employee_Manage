package middlewares

import (
	"employee_manage/config"
	"employee_manage/models"
	"net/http"
	"strconv"
	"strings"

	auth "employee_manage/controllers/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Protect() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.Request.Header.Get("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, auth.MessageResponse{
				Success: false,
				Message: "Token is required",
			})
			c.Abort()
			return
		}

		tokenString = strings.Split(tokenString, " ")[1]

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.ConfigApp.JWT.SecretAccessToken), nil
		})

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			payload := claims["payload"].(map[string]interface{})
			c.Set("payload", payload)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, auth.MessageResponse{
				Success: false,
				Message: "Invalid token",
			})
			c.Abort()
			return
		}
	}
}

func ProtectRole(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		payload := c.GetStringMap("payload")["user_id"].(float64)
		userID := int(payload)

		role, _ := models.GetRole(&user, userID)

		if role == "user" {
			c.JSON(http.StatusForbidden, auth.MessageResponse{
				Success: false,
				Message: "You do not have permission to access this resource",
			})
			c.Abort()
		}

		for _, value := range roles {
			if role == value && role != "manager" {
				c.Next()
				return
			} else if role == value && role == "manager" {
				staffID, _ := strconv.Atoi(c.Param("id"))
				if models.CheckUserInDepartment(userID, staffID) {
					c.Next()
					return
				}
				break
			}
		}

		c.JSON(http.StatusForbidden, auth.MessageResponse{
			Success: false,
			Message: "You do not have permission to access this resource",
		})
		c.Abort()
	}
}
