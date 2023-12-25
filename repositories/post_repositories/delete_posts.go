package post_repositories

import (
	"collagen/config/database"
	"collagen/models/database_model"
	"gorm.io/gorm"
)

func DeletePost(id int) *gorm.DB {
	return database.DB.Delete(database_model.Post{}, id)
}
