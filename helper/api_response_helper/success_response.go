package api_response_helper

import "C"
import (
	"collagen/models/api_response_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateSuccessResponse(c *gin.Context, message string, data any) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	})
}

func GenerateSuccessWithTokenResponse(c *gin.Context, message string, jwt string) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
		"data": api_response_model.WithJwtTokenResponse{
			Token: jwt,
		},
	})
}
