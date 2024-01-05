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

func MigrateAppConfig() error {
	appConfig := []database_model.AppConfig{
		{
			Key:   "maintenance",
			Value: "false",
		},
		{
			Key:   "minimumVersion",
			Value: "146",
		},
		{
			Key:   "operatingTimeWeekday",
			Value: "07.00-20.00 WIB",
		},
		{
			Key:   "customerFriendPhoneNumber",
			Value: "+622150928839",
		},
		{
			Key:   "operatingTimeWeekend",
			Value: "07.00-20.00 WIB",
		},
	}
	trx := DB.Create(&appConfig)
	if trx.Error != nil {
		return trx.Error
	}
	return nil
}
