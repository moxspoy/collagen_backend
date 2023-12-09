package user_detail_repository

import (
	"flop/config/database"
	"flop/models/database_model"
	"gorm.io/gorm"
)

func UpdateUserDetail(request *database_model.UserDetail) *gorm.DB {
	return database.DB.Save(&request)
}
