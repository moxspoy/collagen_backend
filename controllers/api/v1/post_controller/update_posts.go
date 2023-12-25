package post_controller

import (
	"collagen/helper/api_response_helper"
	"collagen/repositories/post_repositories"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UpdatePost handles the request to update an existing post
// @Summary Update an existing post
// @Description Update an existing post
// @Tags Post
// @Accept json
// @Produce json
// @Success		200	{object} api_response_model.SuccessAPIResponse
// @Router			/posts/{id} [put]
// @Param			api_key	header string	true "Api Key"
// @Param post body models.Post true "Updated post object"
func UpdatePost(c *gin.Context) {
	postID, _ := strconv.Atoi(c.Param("id"))

	// Delete the post from the database using GORM
	if err := post_repositories.DeletePost(postID).Error; err != nil {
		api_response_helper.GenerateErrorResponse(c, err)
		return
	}

	// Return a success message in the response
	api_response_helper.GenerateSuccessResponse(c, "Post deleted successfully", postID)
}
