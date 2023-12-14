package user_detail_repository

import (
	"collagen/config/database"
	"collagen/models/database_model"
	"gorm.io/gorm"
)

func GetUserDetail(request *database_model.UserDetail) *gorm.DB {
	return database.DB.First(&request)
}
