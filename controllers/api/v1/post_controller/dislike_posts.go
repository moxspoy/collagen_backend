package post_controller

import (
	"collagen/helper/api_response_helper"
	"collagen/repositories/post_repositories"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

// DislikePost handles the request to dislike a post
// @Summary Dislike a post
// @Description Dislike a post by decrementing the number of likes
//
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Success		200	{object} api_response_model.SuccessAPIResponse
//	@Router			/posts/dislike [post]
//
//	@Param			api_key	header string	true "Api Key"
//
// @Param id path int true "Post ID"
func DislikePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		api_response_helper.GenerateErrorResponse(c, errors.New("invalid post ID"))
		return
	}

	err = post_repositories.DislikePosts(postID)
	if err != nil {
		api_response_helper.GenerateErrorResponse(c, err)
		return
	}

	// Return the posts in the response
	api_response_helper.GenerateSuccessResponse(c, "Like success", postID)
}
