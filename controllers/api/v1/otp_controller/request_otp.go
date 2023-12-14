package otp_controller

import (
	"collagen/helper/api_response_helper"
	"collagen/helper/security_helper"
	"collagen/middleware"
	"collagen/models/database_model"
	"collagen/repositories/one_time_password_repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// RequestOTP godoc
//
//	@Summary		Request OTP
//	@Description	Usually this endpoint used to authenticate user when doing some sensitive data changes
//	@Tags			OTP
//	@Produce		json
//	@Success		200	{object}	api_response_model.SuccessAPIResponse
//	@Router			/otp/request [put]
//	@Param			api_key	header string	true "Api Key"
//	@Security		ApiKeyAuth
func RequestOTP(c *gin.Context) {
	userId := middleware.GetUserIdFromJWT(c)

	now := time.Now()
	then := now.Add(time.Duration(5) * time.Minute)
	token, errOtp := security_helper.GenerateOTP()

	if errOtp != nil {
		c.IndentedJSON(http.StatusInternalServerError, errOtp)
		return
	}

	otp := database_model.OneTimePassword{
		UserID:    userId,
		Token:     token,
		ExpiredAt: then,
		Status:    0,
	}

	one_time_password_repository.RequestOneTimePassword(&otp)

	api_response_helper.GenerateSuccessResponse(c, "OTP has been sent", gin.H{
		"expired_at": then,
	})
}
