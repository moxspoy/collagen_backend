package one_time_password_repository

import (
	"collagen/config/database"
	"collagen/models/database_model"
	"gorm.io/gorm"
)

func RequestOneTimePassword(otp *database_model.OneTimePassword) *gorm.DB {
	return database.DB.Create(&otp)
}
