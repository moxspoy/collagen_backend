package user_repository

import (
	"collagen/config/database"
	"collagen/models/database_model"
	"gorm.io/gorm"
)

func UpdatePhoneNumber(id uint, phoneNumber string) *gorm.DB {
	return database.DB.Model(database_model.User{}).Where("id = ?", id).UpdateColumn("phone_number", phoneNumber)
}
