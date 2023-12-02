package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidateIdentity(c *gin.Context) {
	credential := c.PostForm("credential")
	deviceIdentifier := c.PostForm("device_identifier")
	platform := c.PostForm("platform")
	deviceModel := c.PostForm("device_model")
	version := c.PostForm("version")
	osVersion := c.PostForm("os_version")
	requestId := c.PostForm("request_id")

	// generateJWT

	c.IndentedJSON(http.StatusOK, gin.H{
		"credential":       credential,
		"deviceIdentifier": deviceIdentifier,
		"platform":         platform,
		"deviceModel":      deviceModel,
		"version":          version,
		"osVersion":        osVersion,
		"requestId":        requestId,
	})

}
