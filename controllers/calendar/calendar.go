package calendar

import (
	"employee_manage/models"
	"employee_manage/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Checkin godoc
// @Tags calendar
// @Summary Checkin
// @Description Checkin
// @Accept  json
// @Produce  json
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /calendars/checkin [POST]
func CheckIn(c *gin.Context) {
	userID := utils.GetUserIDByContext(c)

	calendar := models.Calendar{
		UserID: userID,
	}

	err := models.Checkin(&calendar)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "Checkin Success",
	})

}

// Checkout godoc
// @Tags calendar
// @Summary Checkout
// @Description Checkout
// @Accept  json
// @Produce  json
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /calendars/checkout [POST]
func Checkout(c *gin.Context) {
	userID := utils.GetUserIDByContext(c)

	id, err := models.CheckValidCheckout(userID)

	if err != nil {
		_ = c.Error(err)
		return
	}

	timeNow := time.Now().Format("2006-01-02 15:04:05.000")
	calMap := map[string]interface{}{
		"checkout_time": timeNow,
	}

	if _, err := models.Checkout(id, calMap); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "Checkout Success",
	})

}
