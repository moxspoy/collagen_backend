package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func GuardApiKey() gin.HandlerFunc {
	requiredApiKey := os.Getenv("API_KEY")
	if requiredApiKey == "" {
		log.Fatal("API_KEY should be defined in your env")
	}
	return func(context *gin.Context) {
		apiKey := context.Request.Header.Get("API_KEY")
		if apiKey == "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "api key should be defined",
			})
			return
		}
		if apiKey != requiredApiKey {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "api key is wrong",
			})
			return
		}
		context.Next()
	}
}
