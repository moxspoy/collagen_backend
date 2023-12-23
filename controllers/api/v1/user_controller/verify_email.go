package user_controller

import (
	"collagen/helper/api_response_helper"
	"collagen/middleware"
	"collagen/repositories/one_time_password_repository"
	"collagen/repositories/user_repository"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// VerifyEmail godoc
//
//	@Summary		Verify user's email
//	@Description	Usually this endpoint used as part of onboarding. Note that user need to request otp first.
//	@Tags			User
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	api_response_model.SuccessAPIResponse
//	@Router			/user/verify-email [put]
//	@Param			api_key	header string	true "Api Key"
//	@Param			email formData string	true "Email that will be verified"
//	@Param			otp formData string	false "OTP for authentication (if pin already exist)"
//	@Security		ApiKeyAuth
func VerifyEmail(c *gin.Context) {
	currentUserId := middleware.GetUserIdFromJWT(c)
	email := c.Request.FormValue("email")
	otp := c.Request.FormValue("otp")

	if email == "" {
		api_response_helper.GenerateErrorResponse(c, errors.New("email can not be empty"))
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	errs := validate.Var(email, "required,email")
	if errs != nil {
		api_response_helper.GenerateErrorResponse(c, errs)
		return
	}

	// Check OTP
	errs = validate.Var(otp, "required,numeric")
	if errs != nil {
		api_response_helper.GenerateErrorResponse(c, errs)
		return
	}

	errs = one_time_password_repository.CheckOneTimePassword(currentUserId, otp)

	if errs != nil {
		api_response_helper.GenerateErrorResponse(c, errs)
		return
	}

	if tx := user_repository.UpdateEmailVerificationStatus(currentUserId, 1); tx.Error != nil {
		api_response_helper.GenerateErrorResponse(c, tx.Error)
		return
	}

	api_response_helper.GenerateSuccessResponse(c, "Update email verification status successful", email)
}
