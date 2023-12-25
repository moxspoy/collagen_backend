package post_controller

import (
	"collagen/config/database"
	"collagen/helper/api_response_helper"
	"collagen/models/database_model"
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
// @Success	200	{object} api_response_model.SuccessAPIResponse
// @Router /posts/{id} [put]
// @Param api_key	header string	true "Api Key"
// @Param post body models.Post true "Updated post object"
// @Param id path int true "Post ID"
func UpdatePost(c *gin.Context) {
	postID, _ := strconv.Atoi(c.Param("id"))
	var updatedPost database_model.Post

	// Check if the post exists
	if err := database.DB.First(&updatedPost, postID).Error; err != nil {
		api_response_helper.GenerateErrorResponse(c, err)
		return
	}

	// Bind the request body to the updated Post struct
	if err := c.BindJSON(&updatedPost); err != nil {
		api_response_helper.GenerateErrorResponse(c, err)
		return
	}

	// Update the post in the database using GORM
	post_repositories.UpdatePost(updatedPost)

	// Return a success message in the response
	api_response_helper.GenerateSuccessResponse(c, "Post updated successfully", postID)
}
