package database_model

import "gorm.io/gorm"

type AppConfig struct {
	gorm.Model
	Id         uint `json:"id" gorm:"primary_key"`
	Key        string
	Value      string
	LastUpdate string
}

func (AppConfig) TableName() string {
	return "app_config"
}

type AppConfigResponse struct {
	CustomerFriendPhoneNumber string
	Maintenance               string
	MinimumVersion            string
	OperatingTimeWeekday      string
	OperatingTimeWeekend      string
}
