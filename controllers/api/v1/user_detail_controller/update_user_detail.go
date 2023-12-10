package user_detail_controller

import (
	"flop/helper/api_response_helper"
	"flop/middleware"
	"flop/models/database_model"
	"flop/repositories/user_detail_repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// UpdateUserDetail godoc
//
//	@Summary		Update user's detail
//	@Description	Usually this endpoint used to process kyc
//	@Tags			User
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	api_response_model.SuccessAPIResponse
//	@Router			/user-detail/update [put]
//	@Param			api_key	header string	true "Api Key"
//	@Param			data formData database_model.UserDetail	true "User detail that will be saved to the database_model"
//	@Security		ApiKeyAuth
func UpdateUserDetail(c *gin.Context) {
	userId := middleware.GetUserIdFromJWT(c)
	var request database_model.UserDetail
	err := c.Bind(&request)
	request.UserID = userId
	log.Print("wkwk", request.Gender)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	user_detail_repository.UpdateUserDetail(&request)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	api_response_helper.GenerateSuccessResponse(c, "Update user's detail successful", request)
}
