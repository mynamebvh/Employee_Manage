package utils

import "github.com/gin-gonic/gin"

func GetUserIDByContext(c *gin.Context) (userID int) {
	payload := c.GetStringMap("payload")["user_id"].(float64)

	userID = int(payload)
	return
}
