package middlewares

import (
	"employee_manage/config"
	"fmt"
	"net/http"
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
			return
		}

		tokenString = strings.Split(tokenString, " ")[1]

		fmt.Println("Hello", tokenString)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.ConfigApp.JwtConfig.SecretAccessToken), nil
		})

		fmt.Println(err)
		if token == nil {
			c.JSON(http.StatusUnauthorized, auth.MessageResponse{
				Success: false,
				Message: "Invalid token",
			})
			return
		} else {
			c.Next()
		}
	}
}
