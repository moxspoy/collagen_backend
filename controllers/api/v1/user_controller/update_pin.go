package user_controller

import (
	"errors"
	"flop/helper/api_response_helper"
	"flop/helper/security_helper"
	"flop/middleware"
	"flop/repositories/user_repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// UpdatePin godoc
//
//	@Summary		Update user's pin
//	@Description	Usually this endpoint used as part of onboarding.
//	@Tags			User
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	api_response_model.SuccessAPIResponse
//	@Router			/user/update-pin [put]
//	@Param			api_key	header string	true "Api Key"
//	@Param			pin formData string	true "6 digits that will be saved to the database_model"
//	@Security		ApiKeyAuth
func UpdatePin(c *gin.Context) {
	currentUserId := middleware.GetUserIdFromJWT(c)
	pin := c.Request.FormValue("pin")

	if pin == "" {
		api_response_helper.GenerateErrorResponse(c, errors.New("pin can not be empty"))
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	errs := validate.Var(pin, "required,numeric,len=6")

	if errs != nil {
		api_response_helper.GenerateErrorResponse(c, errs)
		return
	}

	hashedPin, hashingError := security_helper.HashPassword(pin)

	if hashingError != nil {
		api_response_helper.GenerateErrorResponse(c, hashingError)
		return
	}

	if dbc := user_repository.UpdatePin(currentUserId, hashedPin); dbc.Error != nil {
		api_response_helper.GenerateErrorResponse(c, dbc.Error)
		return
	}

	api_response_helper.GenerateSuccessResponse(c, "Update pin successful", "")
}
