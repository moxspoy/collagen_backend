package public_controller

import (
	"collagen/config/database"
	"collagen/models/database_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAppInfo godoc
//
//	@Summary		Show application info metadata as the startup info while client app is launched
//	@Description	This endpoint does not require token (public)
//	@Tags			App Config
//	@Accept			json
//	@Produce		json
//	@Success		200	{object} database_model.AppConfigResponse
//	@Router			/app-info [get]
//	@Param			api_key	header string	true "Api Key"
func GetAppInfo(c *gin.Context) {
	var appConfig []database_model.AppConfig
	database.DB.Find(&appConfig)

	response := make(map[string]interface{})

	for _, config := range appConfig {
		response[config.Key] = config.Value
	}
	c.JSON(http.StatusOK, response)
}
