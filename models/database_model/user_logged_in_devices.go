package database_model

import (
	"gorm.io/gorm"
	"time"
)

type UserLoggedInDevices struct {
	gorm.Model
	UserId           uint      `json:"user_id"`
	DeviceModel      string    `json:"device_model"`
	DeviceIdentifier string    `json:"device_identifier"`
	OsVersion        string    `json:"os_version"`
	LastLogin        time.Time `json:"last_login"  gorm:"datetime"`
	Platform         string    `json:"platform"`
	AppNameVersion   string    `json:"app_name_version"`
	AppBuildVersion  int       `json:"app_build_version"`
}

func (UserLoggedInDevices) TableName() string {
	return "user_logged_in_devices"
}
