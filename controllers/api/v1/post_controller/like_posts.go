package post_controller

import (
	"collagen/helper/api_response_helper"
	"collagen/repositories/post_repositories"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

// LikePost handles the request to like a post
// @Summary Like a post
// @Description Like a post by incrementing the number of likes
//
//	@Tags			Post
//	@Accept			json
//	@Produce		json
//	@Success		200	{object} api_response_model.SuccessAPIResponse
//	@Router			/posts/like [post]
//
//	@Param			api_key	header string	true "Api Key"
//
// @Param id path int true "Post ID"
func LikePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		api_response_helper.GenerateErrorResponse(c, errors.New("invalid post ID"))
		return
	}

	err = post_repositories.LikePosts(postID)
	if err != nil {
		api_response_helper.GenerateErrorResponse(c, err)
		return
	}

	// Return the posts in the response
	api_response_helper.GenerateSuccessResponse(c, "Like success", postID)
}
