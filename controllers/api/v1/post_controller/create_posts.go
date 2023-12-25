package post_controller

import (
	"collagen/helper/api_response_helper"
	"collagen/models/database_model"
	"collagen/repositories/post_repositories"
	"github.com/gin-gonic/gin"
)

// CreatePost handles the request to create a new post
// @Summary Create a new post
// @Description Create a new post
//
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Success		200	{object} api_response_model.SuccessAPIResponse
//	@Router			/posts [post]
//	@Param			api_key	header string	true "Api Key"
//
// @Param post body database_model.Post true "Post object"
func CreatePost(c *gin.Context) {
	var newPost database_model.Post

	// Bind the request body to the Post struct
	if err := c.BindJSON(&newPost); err != nil {
		api_response_helper.GenerateErrorResponse(c, err)
		return
	}

	// Create the post in the database using GORM
	post_repositories.CreatePost(newPost)

	// Return the posts in the response
	api_response_helper.GenerateSuccessResponse(c, "", newPost)
}
