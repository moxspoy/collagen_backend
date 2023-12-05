package user_controller

import (
	"errors"
	"flop/helper/api_response_helper"
	"flop/middleware"
	"flop/repositories/users_repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// UpdateUserName godoc
//
//	@Summary		Update user's name
//	@Description	Usually this endpoint used as part of onboarding.
//	@Tags			User
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	api_response_model.SuccessAPIResponse
//	@Router			/user/update-user-name [put]
//	@Param			api_key	header string	true "Api Key"
//	@Param			name formData string	true "Name that will be saved to the database_model"
//	@Security		ApiKeyAuth
func UpdateUserName(c *gin.Context) {
	currentUserId := middleware.GetUserIdFromJWT(c)
	name := c.Request.FormValue("name")

	if name == "" {
		api_response_helper.GenerateErrorResponse(c, errors.New("name can not be empty"))
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	errs := validate.Var(name, "required,alphanumeric")

	if errs != nil {
		api_response_helper.GenerateErrorResponse(c, errs)
		return
	}

	if dbc := users_repository.UpdateUserName(currentUserId, name); dbc.Error != nil {
		api_response_helper.GenerateErrorResponse(c, dbc.Error)
		return
	}

	api_response_helper.GenerateSuccessResponse(c, "Update name successful", name)
}
