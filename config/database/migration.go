package database

import (
	"flop/models/database_model"
	"gorm.io/gorm"
)

func RunAutoMigration(database *gorm.DB) error {
	return database.AutoMigrate(
		&database_model.User{},
		&database_model.AppConfig{},
		&database_model.UserLoggedInDevices{},
		&database_model.OneTimePassword{},
		&database_model.UserDetail{},
	)
}
