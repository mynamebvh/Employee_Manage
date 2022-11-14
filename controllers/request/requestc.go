package request

import (
	errorModels "employee_manage/constant"
	"employee_manage/models"
	"employee_manage/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetRequests godoc
// @Tags request
// @Summary Get list request
// @Description Get list request
// @Success 200 {object} PaginationResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Security Authentication
// @Router /requests [GET]
func GetRequests(c *gin.Context) {
	role := c.GetStringMap("role")

	result, err := models.GetRequests(c, role)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, PaginationResponse{
		Success:  true,
		Message:  "Success",
		Data:     result.Data,
		Current:  result.Current,
		PageSize: result.PageSize,
		Total:    int(result.Total),
	})
}

// GetRequestByID godoc
// @Tags request
// @Summary Get request by id
// @Description Get request by id
// @Param request_id path int true "id of request"
// @Success 200 {object} MessageResponse
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Security Authentication
// @Router /requests/{request_id} [GET]
func GetRequestByID(c *gin.Context) {

	requestID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		appError := errorModels.NewAppError(errors.New("request id is invalid"), errorModels.ValidationError)
		c.Error(appError)
		return
	}

	request, err := models.GetRequestByID(requestID)

	if err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		c.Error(appError)
		return
	}

	if request.ID == 0 {
		appError := errorModels.NewAppError(errors.New("request not found"), errorModels.ValidationError)
		c.Error(appError)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Message: "Success",
		Data:    request,
	})
}

// CreateRequest godoc
// @Tags request
// @Summary Create new request
// @Description Create new request
// @Accept  json
// @Produce  json
// @Param data body NewRequest true "body data"
// @Success 200 {object} models.Request
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Security Authentication
// @Router /requests [POST]
func CreateRequest(c *gin.Context) {
	var newRequest NewRequest

	if err := c.BindJSON(&newRequest); err != nil {
		appError := errorModels.NewAppError(err, errorModels.ValidationError)
		_ = c.Error(appError)
		return
	}

	errValidate := validateStructRequest(newRequest)

	if errValidate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   errValidate,
		})
		return
	}

	userID := utils.GetUserIDByContext(c)

	request := models.Request{
		Type:    newRequest.Type,
		Content: newRequest.Content,
		UserID:  userID,
		Status:  "pending",
	}

	err := models.CreateRequest(&request)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, MessageResponse{
		Success: true,
		Message: "Success",
		Data:    request,
	})
}

// UpdateRequestByID godoc
// @Tags request
// @Summary Update request by ID
// @Description Update request by ID on the system
// @Param request_id path int true "id of request"
// @Param data body NewRequest true "body data"
// @Success 200 {object} models.Request
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Security Authentication
// @Router /requests/{request_id} [PUT]
func UpdateRequestByID(c *gin.Context) {
	requestID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appError := errorModels.NewAppError(errors.New("department id is invalid"), errorModels.ValidationError)
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

	userID := utils.GetUserIDByContext(c)
	requestMap["approved_by"] = userID

	request, err := models.UpdateRequestByID(requestID, requestMap)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, MessageResponse{
		Success: true,
		Data:    request,
	})

}
