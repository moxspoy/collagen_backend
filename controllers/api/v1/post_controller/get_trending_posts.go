package post_controller

import (
	"collagen/helper/api_response_helper"
	"collagen/repositories/post_repositories"
	"github.com/gin-gonic/gin"
	"time"
)

// GetTrendingPosts handles the request to retrieve trending posts
// @Summary Get trending posts
// @Description Get posts that are currently trending based on likes, comments, and shares within a time frame
// @Accept json
// @Produce json
// @Param duration query string false "Time duration for trending calculation (e.g., 24h)"
// @Success 200 {object} api_response_model.SuccessAPIResponse
// @Router /post/trending [get]
func GetTrendingPosts(c *gin.Context) {
	// Parse query parameter for the duration of trending calculation
	durationString := c.DefaultQuery("duration", "24h")
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		api_response_helper.GenerateErrorResponse(c, err)
		return
	}

	// Calculate the cutoff time based on the duration
	cutoffTime := time.Now().Add(-duration)

	trendingPosts := post_repositories.GetTrendingPosts(cutoffTime)

	// Return the trending posts in the response
	api_response_helper.GenerateSuccessResponse(c, "", trendingPosts)
}
