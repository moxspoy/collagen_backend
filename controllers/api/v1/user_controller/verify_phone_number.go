package user_controller

import (
	"errors"
	"flop/helper/api_response_helper"
	"flop/middleware"
	"flop/repositories/one_time_password_repository"
	"flop/repositories/user_repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// VerifyPhoneNumber godoc
//
//	@Summary		Verify user's phone number
//	@Description	Usually this endpoint used as part of onboarding. Note that user need to request otp first.
//	@Tags			User
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	api_response_model.SuccessAPIResponse
//	@Router			/user/verify-phone-number [put]
//	@Param			api_key	header string	true "Api Key"
//	@Param			phone_number formData string	true "Phone number that will be verified"
//	@Param			otp formData string	false "OTP for authentication (if pin already exist)"
//	@Security		ApiKeyAuth
func VerifyPhoneNumber(c *gin.Context) {
	currentUserId := middleware.GetUserIdFromJWT(c)
	phoneNumber := c.Request.FormValue("phone_number")
	otp := c.Request.FormValue("otp")

	if phoneNumber == "" {
		api_response_helper.GenerateErrorResponse(c, errors.New("phone number can not be empty"))
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	errs := validate.Var(phoneNumber, "required,e164")
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

	if tx := user_repository.UpdatePhoneVerificationStatus(currentUserId, 1); tx.Error != nil {
		api_response_helper.GenerateErrorResponse(c, tx.Error)
		return
	}

	api_response_helper.GenerateSuccessResponse(c, "Update phone verification status successful", phoneNumber)
}
