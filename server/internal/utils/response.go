package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Items    []interface{} `json:"items"`
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
}

func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})

}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Success: false,
		Error:   message,
	})
}

func PaginatedSuccessResponse(c *gin.Context, statusCode int, data interface{}, total uint64, pageSizes int, page int) {
	c.JSON(statusCode, PaginatedResponse{
		Items:    data.([]interface{}),
		Total:    int(total),
		Page:     page,
		PageSize: pageSizes,
	})
}
