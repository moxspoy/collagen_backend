package user_controller

import (
	"collagen/helper/api_response_helper"
	"collagen/middleware"
	"collagen/models/database_model"
	"collagen/repositories/one_time_password_repository"
	"collagen/repositories/user_repository"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// UpdateEmail godoc
//
//	@Summary		Update user's email
//	@Description	Usually this endpoint used because user fill phone number first. Note that user need to request otp first
//	@Tags			User
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	api_response_model.SuccessAPIResponse
//	@Router			/user/update-email [put]
//	@Param			api_key	header string	true "Api Key"
//	@Param			newEmail formData string	true "Email that will be saved to the database_model"
//	@Param			otp formData string	false "OTP for authentication (if pin already exist)"
//	@Security		ApiKeyAuth
func UpdateEmail(c *gin.Context, authMiddleware *jwt.GinJWTMiddleware) {
	userId := middleware.GetUserIdFromJWT(c)
	newEmail := c.Request.FormValue("newEmail")
	otp := c.Request.FormValue("otp")

	// Validation
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Var(newEmail, "required,email")

	if err != nil {
		api_response_helper.GenerateErrorResponse(c, err)
		return
	}

	// Check OTP
	user := user_repository.GetOneUserById(userId)
	if user.Email.Valid && user.IsEmailVerified() {
		err = validate.Var(otp, "required,numeric")

		if err != nil {
			api_response_helper.GenerateErrorResponse(c, err)
			return
		}

		err = one_time_password_repository.CheckOneTimePassword(userId, otp)
		if err != nil {
			api_response_helper.GenerateErrorResponse(c, err)
			return
		}
	}

	// OTP is match, update email and JWT
	user = database_model.User{}
	user_repository.UpdateUserEmail(&user, userId, newEmail)

	newJWT, _, err := authMiddleware.TokenGenerator(&user)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	api_response_helper.GenerateSuccessWithTokenResponse(c, "Update email successful", newJWT)
}
