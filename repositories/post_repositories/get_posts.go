package post_repositories

import (
	"collagen/config/database"
	"collagen/models/database_model"
)

// GetPosts handles the request to retrieve posts with pagination
func GetPosts(offset int, limit int) []database_model.Post {

	// Fetch posts from the database using GORM
	var posts []database_model.Post
	database.DB.Offset(offset).Limit(limit).Find(&posts)
	return posts
}
