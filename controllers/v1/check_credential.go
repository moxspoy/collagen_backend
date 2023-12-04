package v1

import (
	"flop/config/database"
	"flop/models"
	"flop/models/api_responses"
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
//	@Tags			Auth
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	api_responses.CheckCredentialResponse
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
	isUserExist := len(users) > 0

	user := models.Users{}
	if isUserExist {
		user = users[0]
	}
	email := user.Email
	phone := user.PhoneNumber

	response := api_responses.CheckCredentialResponse{
		Email:           email,
		PhoneNumber:     phone,
		IsUserExist:     isUserExist,
		IsEmailVerified: user.IsPhoneVerified(),
		IsPhoneVerified: user.IsPhoneVerified(),
		IsPinRegistered: user.IsPinRegistered(),
		IsRegistered:    true,
	}

	c.IndentedJSON(http.StatusOK, response)
}
