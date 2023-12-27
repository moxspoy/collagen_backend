package post_repositories

import (
	"collagen/config/database"
	"collagen/models/database_model"
	"gorm.io/gorm"
	"time"
)

func GetTrendingPosts(cutoffTime time.Time) *gorm.DB {
	// Fetch trending posts from the database using GORM
	var trendingPosts []database_model.Post
	return database.DB.Where("created_at >= ?", cutoffTime).
		Order("num_likes + num_comments + num_shares DESC").
		Limit(10).Find(&trendingPosts)

}
