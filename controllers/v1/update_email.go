package v1

import (
	"flop/helper/api_response"
	"flop/middleware"
	"flop/models/database_model"
	"flop/repositories/users"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateEmail godoc
//
//	@Summary		Update user's email
//	@Description	Usually this endpoint used because user fill phone number first
//	@Tags			User
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	api_response_model.SuccessAPIResponse
//	@Router			/user/update-email [put]
//	@Param			api_key	header string	true "Api Key"
//	@Param			new_email formData string	true "Email that will be saved to the database_model"
//	@Security		ApiKeyAuth
func UpdateEmail(c *gin.Context, authMiddleware *jwt.GinJWTMiddleware) {
	currentEmail := middleware.GetEmailFromJWT(c)
	newEmail := c.Request.FormValue("new_email")

	user := database_model.Users{}
	users.UpdateUserEmail(&user, currentEmail, newEmail)

	newJWT, _, err := authMiddleware.TokenGenerator(&user)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	api_response.GenerateSuccessResponse(c, "Update email successful", newJWT)
}
