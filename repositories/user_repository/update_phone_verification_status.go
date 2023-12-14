package user_repository

import (
	"flop/config/database"
	"flop/models/database_model"
	"gorm.io/gorm"
)

func UpdatePhoneVerificationStatus(id uint, status int) *gorm.DB {
	return database.DB.Model(database_model.User{}).Where("id = ?", id).UpdateColumn("phone_verification_status", status)
}
