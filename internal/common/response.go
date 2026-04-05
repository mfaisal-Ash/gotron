package common

import "github.com/gin-gonic/gin"

type Meta struct {
	page    int         `json:"page"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
	TotalPage int         `json:"total_page"`
}

type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

func JSONSuccessResponse(c *gin.Context, status int, message string, data interface{}, meta *Meta) {
	c.JSON(status, SuccessResponse{
		Success:  true,
		Status:  status,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func JSONErrorResponse(c *gin.Context, status int, message string, errors map[string]interface{}) {
	c.JSON(status, ErrorResponse{
		Success:  false,
		Status:  status,
		Message: message,
		Errors: errors,
	})
}