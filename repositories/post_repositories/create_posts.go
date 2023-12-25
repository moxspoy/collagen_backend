package post_repositories

import (
	"collagen/config/database"
	"collagen/models/database_model"
	"gorm.io/gorm"
)

func CreatePost(post database_model.Post) *gorm.DB {
	return database.DB.Create(&post)
}
