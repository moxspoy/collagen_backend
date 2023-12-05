package auth_controller

import (
	"flop/helper/api_response_helper"
	"flop/models/api_request_model"
	"flop/models/database_model"
	"flop/repositories/user_logged_in_devices_repository"
	"flop/repositories/users_repository"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserSignUp godoc
//
//	@Summary		Sign up endpoint
//	@Description	This endpoint used to register user to the platform
//	@Tags			Auth
//	@Produce		json
//	@Accept			multipart/form-data
//	@Success		200	{object} api_response_model.SuccessAPIResponse
//	@Router			/auth/sign-up [post]
//	@Param			api_key	header string	true "Api Key"
//	@Param			data formData api_request_model.UserSignUpRequest true "Request object"
func UserSignUp(c *gin.Context, authMiddleware *jwt.GinJWTMiddleware) {
	var request api_request_model.UserSignUpRequest
	err := c.Bind(&request)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// save user
	user := database_model.Users{
		Email:                   request.Email,
		PhoneNumber:             request.PhoneNumber,
		Name:                    request.Name,
		Password:                "",
		Pin:                     "",
		Status:                  0,
		PhoneVerificationStatus: 0,
		EmailVerificationStatus: 0,
	}
	users_repository.InsertUser(&user)

	// save device identifier
	userLoggedInDevice := database_model.UserLoggedInDevices{
		UserId:           user.ID,
		DeviceModel:      request.DeviceModel,
		DeviceIdentifier: request.DeviceIdentifier,
		OsVersion:        request.OsVersion,
		Platform:         request.Platform,
		AppNameVersion:   request.AppNameVersion,
		AppBuildVersion:  request.AppBuildVersion,
	}

	user_logged_in_devices_repository.InsertOnSignUp(&userLoggedInDevice)

	newJWT, _, err := authMiddleware.TokenGenerator(&user)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	api_response_helper.GenerateSuccessResponse(c, "success register", newJWT)
}
