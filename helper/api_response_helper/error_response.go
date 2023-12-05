package api_response_helper

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GenerateErrorResponse(c *gin.Context, err error) {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{
		"code":    http.StatusInternalServerError,
		"message": err.Error(),
		"data":    err,
	})
}
