package post_controller

import (
	"collagen/helper/api_response_helper"
	"collagen/repositories/post_repositories"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetPosts godoc
//
// @Summary Get posts with pagination
// @Description Get posts with pagination
//
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Success		200	{object} api_response_model.SuccessAPIResponse
//	@Router			/posts/ [get]
//
// @Param page query int false "Page number"
// @Param limit query int false "Number of items per page"
//
//	@Param			api_key	header string	true "Api Key"
func GetPosts(c *gin.Context) {
	// Parse query parameters for pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Calculate the offset based on the page and limit
	offset := (page - 1) * limit

	// Fetch posts from the database using GORM
	posts := post_repositories.GetPosts(offset, limit)

	// Return the posts in the response
	api_response_helper.GenerateSuccessResponse(c, "", posts)
}
