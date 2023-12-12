package user_detail_controller

import (
	"flop/helper/upload_helper"
	"github.com/gin-gonic/gin"
)

// UpdateUserIdentityImage godoc
//
//	@Summary		Update user's identify image
//	@Description	Usually this endpoint used to process kyc
//	@Tags			User
//	@Accept			multipart/form-data
//	@Produce		json
//	@Success		200	{object}	api_response_model.SuccessAPIResponse
//	@Router			/user-detail/update [put]
//	@Param			api_key	header string	true "Api Key"
//	@Param			file formData file	true "User image that will be saved to the database_model"
//	@Security		ApiKeyAuth
func UpdateUserIdentityImage(c *gin.Context) {
	upload_helper.UploadImageForKyc(c, "identity_images/")
}
