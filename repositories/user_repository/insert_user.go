package user_repository

import (
	"flop/config/database"
	"flop/models/database_model"
	"gorm.io/gorm"
)

func InsertUser(user *database_model.User) *gorm.DB {
	return database.DB.Create(&user)
}
