package universities_controller

import (
	"collagen/helper/api_response_helper"
	"collagen/models/database_model"
	"collagen/repositories/university_repositories"
	"github.com/gin-gonic/gin"
)

// CreateUniversity godoc
//
//	@Summary		Create university object.
//	@Description	This endpoint used to create university when it is not available yet on the database
//	@Tags			User
//	@Produce		json
//	@Success		200	{object} api_response_model.SuccessAPIResponse
//	@Router			/universities/create [post]
//	@Param			api_key	header string	true "Api Key"
//	@Security		ApiKeyAuth
//	@Param			data formData database_model.University true "Request object"
func CreateUniversity(c *gin.Context) {
	var request database_model.University
	if err := c.ShouldBind(&request); err != nil {
		api_response_helper.GenerateErrorResponse(c, err)
		return
	}
	trx := university_repositories.CreateUniversity(request)
	if trx.Error != nil {
		api_response_helper.GenerateErrorResponse(c, trx.Error)
		return
	}
	api_response_helper.GenerateSuccessResponse(c, "Create university success", request)
}
