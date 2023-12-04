package user_logged_in_devices

import (
	"flop/config/database"
	database2 "flop/models/database_model"
	"gorm.io/gorm"
)

func InsertOnSignUp(request *database2.UserLoggedInDevices) *gorm.DB {
	return database.DB.Create(&request)
}
