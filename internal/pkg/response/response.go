package response

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Success bool           `json:"success"`
	Data    interface{}    `json:"data,omitempty"`
	Meta    interface{}    `json:"meta,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
}

// Success sends a successful response with data
func Success(c *gin.Context, data interface{}) {
	c.JSON(200, &Response{
		Success: true,
		Data:    data,
	})
}

// SuccessWithPagination sends a successful response with pagination metadata
func SuccessWithPagination(c *gin.Context, data interface{}, total int64, page, pageSize int) {
	c.JSON(200, &Response{
		Success: true,
		Data:    data,
		Meta: map[string]interface{}{
			"pagination": map[string]interface{}{
				"total":     total,
				"page":      page,
				"page_size": pageSize,
				"pages":     (total + int64(pageSize) - 1) / int64(pageSize),
			},
		},
	})
}

// Error sends an error response
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, &Response{
		Success: false,
		Error: &ErrorResponse{
			Code:    code,
			Message: message,
		},
	})
}

// StringToUint converts string to uint safely
func StringToUint(s string) uint {
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint(val)
}
