package user_repository

import (
	"collagen/config/database"
	"collagen/models/database_model"
	"gorm.io/gorm"
)

func UpdateEmailVerificationStatus(id uint, status int) *gorm.DB {
	return database.DB.Model(database_model.User{}).Where("id = ?", id).UpdateColumn("email_verification_status", status)
}
