package user_logged_in_devices_repository

import (
	"flop/config/database"
	"flop/models/database_model"
	"gorm.io/gorm"
)

func InsertOnSignUp(request *database_model.UserLoggedInDevices) *gorm.DB {
	return database.DB.Create(&request)
}
