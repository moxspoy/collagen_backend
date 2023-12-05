package user_controller

import (
	"flop/helper/api_response_helper"
	"flop/middleware"
	"flop/models/database_model"
	"flop/repositories/users_repository"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// UpdatePhoneNumber godoc
//
//	@Summary		Update user's phone number
//	@Description	Usually this endpoint used as part of onboarding. Note that it should contains country code like +6285911110000
//	@Tags			User
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	api_response_model.SuccessAPIResponse
//	@Router			/user/update-phone-number [put]
//	@Param			api_key	header string	true "Api Key"
//	@Param			phone_number formData string	true "Phone number that will be saved to the database_model"
//	@Security		ApiKeyAuth
func UpdatePhoneNumber(c *gin.Context, authMiddleware *jwt.GinJWTMiddleware) {
	currentUserId := middleware.GetUserIdFromJWT(c)
	phoneNumber := c.Request.FormValue("phone_number")

	user := database_model.Users{
		ID:          currentUserId,
		PhoneNumber: phoneNumber,
	}
	users_repository.UpdatePhoneNumber(&user)

	api_response_helper.GenerateSuccessResponse(c, "Update phone number successful", phoneNumber)
}
