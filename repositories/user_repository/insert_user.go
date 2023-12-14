package user_repository

import (
	"collagen/config/database"
	"collagen/models/database_model"
	"gorm.io/gorm"
)

func InsertUser(user *database_model.User) *gorm.DB {
	return database.DB.Create(&user)
}
