package user_detail_repository

import (
	"flop/config/database"
	"flop/models/database_model"
	"gorm.io/gorm"
)

func GetUserDetail(request *database_model.UserDetail) *gorm.DB {
	return database.DB.First(&request)
}
