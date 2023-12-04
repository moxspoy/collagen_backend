package users

import (
	"flop/config/database"
	database2 "flop/models/database_model"
	"gorm.io/gorm"
)

func InsertUser(user *database2.Users) *gorm.DB {
	return database.DB.Create(&user)
}
