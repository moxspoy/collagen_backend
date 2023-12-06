package users_repository

import (
	"flop/config/database"
	"flop/models/database_model"
	"gorm.io/gorm"
)

func UpdateUserName(id uint, name string) *gorm.DB {
	return database.DB.Model(database_model.Users{}).Where("id = ?", id).UpdateColumn("name", name)
}