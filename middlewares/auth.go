package middlewares

import (
	"employee_manage/config"
	"employee_manage/models"
	"employee_manage/utils"
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

func ProtectRole(roles []string, model string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		userID := utils.GetUserIDByContext(c)

		role, _ := models.GetRole(&user, userID)

		if role.Name == "user" {
			c.JSON(http.StatusForbidden, auth.MessageResponse{
				Success: false,
				Message: "You do not have permission to access this resource",
			})
			c.Abort()
			return
		}

		for _, value := range roles {
			if role.Name == value && role.Name != "manager" {
				c.Set("role", map[string]interface{}{
					"name":          role.Name,
					"department_id": role.DepartmentID,
				})
				c.Next()
				return
			} else if role.Name == value && role.Name == "manager" {
				id, _ := strconv.Atoi(c.Param("id"))

				isNext, err := models.CheckManageAccess(id, role.DepartmentID, model)
				if err != nil {
					c.JSON(http.StatusForbidden, auth.MessageResponse{
						Success: false,
						Message: "You do not have permission to access this resource",
					})
					c.Abort()
					return
				}

				if isNext {
					c.Set("role", map[string]interface{}{
						"name":          role.Name,
						"department_id": role.DepartmentID,
					})
					c.Next()
					return
				} else {
					c.JSON(http.StatusForbidden, auth.MessageResponse{
						Success: false,
						Message: "You do not have permission to access this resource",
					})
					c.Abort()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, auth.MessageResponse{
			Success: false,
			Message: "You do not have permission to access this resource",
		})
		c.Abort()
	}
}
