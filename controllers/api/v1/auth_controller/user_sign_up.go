package auth_controller

import (
	"database/sql"
	"flop/helper/api_response_helper"
	"flop/models/api_request_model"
	"flop/models/database_model"
	"flop/repositories/user_logged_in_devices_repository"
	"flop/repositories/user_repository"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
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
	email := request.Credential
	phoneNumber := request.Credential
	if strings.Contains(request.Credential, "@") {
		phoneNumber = ""
	} else {
		email = ""
	}
	isEmailNull := email == ""
	isPhoneNull := phoneNumber == ""

	user := database_model.User{
		Email:                   sql.NullString{String: email, Valid: !isEmailNull},
		PhoneNumber:             sql.NullString{String: phoneNumber, Valid: !isPhoneNull},
		Name:                    request.Name,
		Password:                "",
		Pin:                     "",
		Status:                  0,
		PhoneVerificationStatus: 0,
		EmailVerificationStatus: 0,
	}
	result := user_repository.InsertUser(&user)
	if result.Error != nil {
		api_response_helper.GenerateErrorResponse(c, result.Error)
		return
	}

	// save device identifier
	userLoggedInDevice := database_model.UserLoggedInDevices{
		UserID:           user.ID,
		DeviceModel:      request.DeviceModel,
		DeviceIdentifier: request.DeviceIdentifier,
		OsVersion:        request.OsVersion,
		Platform:         request.Platform,
		AppNameVersion:   request.AppNameVersion,
		AppBuildVersion:  request.AppBuildVersion,
		LastLogin:        time.Now().Format("2006-01-02 15:04:05"),
	}

	user_logged_in_devices_repository.InsertOnSignUp(&userLoggedInDevice)

	newJWT, _, jwtErr := authMiddleware.TokenGenerator(&user)

	if jwtErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, jwtErr)
		return
	}
	api_response_helper.GenerateSuccessWithTokenResponse(c, "success register", newJWT)
}
