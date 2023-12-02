package v1

import (
	"flop/config/database"
	"flop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAppInfo(c *gin.Context) {
	var appConfig []models.AppConfig
	database.DB.Find(&appConfig)

	response := make(map[string]interface{})

	for _, config := range appConfig {
		response[config.Key] = config.Value
	}
	c.JSON(http.StatusOK, response)
}
