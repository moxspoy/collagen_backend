package universities_controller

import (
	"collagen/helper/api_response_helper"
	"collagen/repositories/university_repositories"
	"github.com/gin-gonic/gin"
)

// GetAllUniversities godoc
//
//	@Summary		Get all universities object
//	@Description	This endpoint used to fetch all universities data
//	@Tags			User
//	@Produce		json
//	@Success		200	{object} api_response_model.SuccessAPIResponse
//	@Router			/universities/all [get]
//	@Param			api_key	header string	true "Api Key"
//	@Security		ApiKeyAuth
func GetAllUniversities(c *gin.Context) {
	universities := university_repositories.GetAllUniversities()
	api_response_helper.GenerateSuccessResponse(c, "", universities)
}
