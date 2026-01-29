package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Response(c *gin.Context, code int, message string, data interface{}) {
	res := JsonResponse{
		Code:    code,
		Message: message,
	}

	if data != nil {
		res.Data = data
	}
	c.JSON(code, res)
	return
}

func SuccessData(c *gin.Context, message string, data interface{}) {
	response(c, 0, message, data)
}

func ErrorData(c *gin.Context, message string, data interface{}) {
	response(c, 1, message, data)
}

func ForbiddenData(c *gin.Context, message string, data interface{}) {
	jsonData := JsonResponse{
		Code:    403,
		Message: message,
	}

	if data != nil {
		jsonData.Data = data
	}

	c.JSON(http.StatusForbidden, jsonData)
	c.Abort()
}

func response(c *gin.Context, code int, message string, data interface{}) {
	jsonData := JsonResponse{
		Code:    code,
		Message: message,
	}

	if data != nil {
		jsonData.Data = data
	}

	c.JSON(http.StatusOK, jsonData)
	c.Abort()
}
