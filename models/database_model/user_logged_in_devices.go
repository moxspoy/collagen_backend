package database_model

import "time"

type UserLoggedInDevices struct {
	Id               uint      `json:"id"`
	UserId           uint      `json:"user_id"`
	DeviceModel      string    `json:"device_model"`
	DeviceIdentifier string    `json:"device_identifier"`
	OsVersion        string    `json:"os_version"`
	CreatedAt        time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	LastLogin        time.Time `json:"last_login"  gorm:"default:CURRENT_TIMESTAMP()"`
	Platform         string    `json:"platform"`
	AppNameVersion   string    `json:"app_name_version"`
	AppBuildVersion  int       `json:"app_build_version"`
}

func (UserLoggedInDevices) TableName() string {
	return "user_logged_in_devices"
}
