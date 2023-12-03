package v1

import (
	"flop/config/database"
	"flop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	EmailCredential = "email"
)

// CheckCredential godoc
//
//	@Summary		Check whether email or phone number exist on the database
//	@Description	Usually this endpoint used before validate user's identity
//	@Tags			check-credential
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	models.AppConfig
//	@Router			/auth/check-credential [post]
//	@Param			api_key	header string	true "Api Key"
//	@Param			credential formData string	true "Email/Phone Number"
func CheckCredential(c *gin.Context) {

	isPhoneNumber := c.Query("type") == EmailCredential
	credential := c.PostForm("credential")

	if credential == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Credential should not be empty",
		})
		return
	}

	var users []models.Users
	var whereClause = "email = ?"
	if isPhoneNumber {
		whereClause = "phone_number = ?"
	}
	database.DB.Where(whereClause, credential).First(&users)

	if len(users) <= 0 {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "User not exist",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "User exist",
	})
}
