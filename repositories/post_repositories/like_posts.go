package post_repositories

import (
	"collagen/config/database"
	"collagen/models/database_model"
)

func LikePosts(postID int) error {
	// Fetch the post from the database using GORM
	var post database_model.Post
	if err := database.DB.First(&post, postID).Error; err != nil {
		return err
	}

	// Increment the number of likes
	post.NumLikes++

	// Update the post in the database using GORM
	database.DB.Save(&post)
	return nil
}
