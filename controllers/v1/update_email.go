package v1

import (
	"flop/helper/api_response"
	"flop/repositories/users"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// UpdateEmail godoc
//
//	@Summary		Update user's email
//	@Description	Usually this endpoint used because user fill phone number first
//	@Tags			User
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	models.SuccessAPIResponseMessageAndCode
//	@Router			/user/update-email [put]
//	@Param			api_key	header string	true "Api Key"
//	@Param			new_email formData string	true "Email that will be saved to the database"
//	@Security		ApiKeyAuth
func UpdateEmail(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	oldEmail := fmt.Sprintf("%v", claims[identityKey])
	newEmail := c.Request.FormValue("new_email")
	users.UpdateUserEmail(oldEmail, newEmail)
	api_response.GenerateSuccessResponse(c, "Update email successful")
}
