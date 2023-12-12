package user_repository

import (
	"flop/config/database"
	"flop/models/database_model"
	"gorm.io/gorm"
)

func UpdatePhoneNumber(id uint, phoneNumber string) *gorm.DB {
	return database.DB.Model(database_model.User{}).Where("id = ?", id).UpdateColumn("phone_number", phoneNumber)
}
