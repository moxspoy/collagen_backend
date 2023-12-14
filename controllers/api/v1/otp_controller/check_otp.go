package otp_controller

import (
	"errors"
	"flop/helper/api_response_helper"
	"flop/middleware"
	"flop/repositories/one_time_password_repository"
	"github.com/gin-gonic/gin"
)

// CheckOTP godoc
//
//	@Summary		Check OTP
//	@Description	Usually this endpoint used to authenticate user when doing some sensitive data changes
//	@Tags			OTP
//	@Produce		json
//	@Accept			multipart/form-data
//	@Success		200	{object}	api_response_model.SuccessAPIResponse
//	@Router			/otp/check [post]
//	@Param			api_key	header string	true "Api Key"
//	@Param			otp formData string	true "OTP from SMS/Whatsapp to be checked"
//	@Security		ApiKeyAuth
func CheckOTP(c *gin.Context) {
	userId := middleware.GetUserIdFromJWT(c)
	otp := c.PostForm("otp")

	if otp == "" {
		api_response_helper.GenerateErrorResponse(c, errors.New("OTP can not be empty"))
		return
	}

	err := one_time_password_repository.CheckOneTimePassword(userId, otp)

	if err != nil {
		api_response_helper.GenerateErrorResponse(c, err)
		return
	}

	api_response_helper.GenerateSuccessResponse(c, "Success", "")
}
