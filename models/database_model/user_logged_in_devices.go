package database_model

import (
	"gorm.io/gorm"
)

type UserLoggedInDevices struct {
	gorm.Model
	UserID           uint
	User             User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DeviceModel      string
	DeviceIdentifier string
	OsVersion        string
	LastLogin        string
	Platform         string
	AppNameVersion   string
	AppBuildVersion  int
}

func (UserLoggedInDevices) TableName() string {
	return "user_logged_in_devices"
}
