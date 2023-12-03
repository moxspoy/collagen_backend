package api_response

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateSuccessResponse(c *gin.Context, message string) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
	})
}
