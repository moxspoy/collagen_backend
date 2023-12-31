package database

import (
	"collagen/models/database_model"
	"gorm.io/gorm"
)

func RunAutoMigration(database *gorm.DB) error {
	return database.AutoMigrate(
		&database_model.User{},
		&database_model.AppConfig{},
		&database_model.UserLoggedInDevices{},
		&database_model.OneTimePassword{},
		&database_model.UserDetail{},
		&database_model.University{},
		&database_model.IndonesiaArea{},
		&database_model.Post{},
		&database_model.Comment{},
	)
}
