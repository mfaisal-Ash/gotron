package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	RequestID string `json:"request_id,omitempty"`
}

type ErrorResponse struct {
	Meta   Meta `json:"meta"`
	Errors any  `json:"errors,omitempty"`
}

type SuccessResponse struct {
	Meta       Meta            `json:"meta"`
	Data       any             `json:"data,omitempty"`
	Pagination *PaginationMeta `json:"pagination,omitempty"`
}

func Success(c *gin.Context, code int, message string, data any) {
	c.JSON(code, SuccessResponse{
		Meta: Meta{
			Code:      code,
			Status:    "success",
			Message:   message,
			RequestID: RequestID(c),
		},
		Data: data,
	})
}

func SuccessWithPagination(c *gin.Context, code int, message string, data any, pagination PaginationMeta) {
	c.JSON(code, SuccessResponse{
		Meta: Meta{
			Code:      code,
			Status:    "success",
			Message:   message,
			RequestID: RequestID(c),
		},
		Data:       data,
		Pagination: &pagination,
	})
}

func Fail(c *gin.Context, err error) {
	appErr := AsAppError(err)
	if appErr == nil {
		appErr = Internal("internal server error", nil)
	}

	c.JSON(appErr.Code, ErrorResponse{
		Meta: Meta{
			Code:      appErr.Code,
			Status:    "error",
			Message:   appErr.Message,
			RequestID: RequestID(c),
		},
		Errors: appErr.Details,
	})
}

func Abort(c *gin.Context, err error) {
	Fail(c, err)
	c.Abort()
}

func RequestID(c *gin.Context) string {
	if c == nil {
		return ""
	}
	value, ok := c.Get(ContextKeyRequestID)
	if !ok {
		return ""
	}
	requestID, ok := value.(string)
	if !ok {
		return ""
	}
	return requestID
}

func AbortMethodNotAllowed(c *gin.Context) {
	Abort(c, NewAppError(http.StatusMethodNotAllowed, "method not allowed", nil, nil))
}
