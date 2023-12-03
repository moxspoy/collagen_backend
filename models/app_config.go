package models

type AppConfig struct {
	Id         uint   `json:"id"`
	Key        string `json:"key"`
	Value      string `json:"value"`
	LastUpdate string `json:"last_update"`
}

func (AppConfig) TableName() string {
	return "app_config"
}

type AppConfigResponse struct {
	CustomerFriendPhoneNumber string `json:"customer_friend_phone_number"`
	Maintenance               string `json:"maintenance"`
	MinimumVersion            string `json:"minimum_version"`
	OperatingTimeWeekday      string `json:"operating_time_weekday"`
	OperatingTimeWeekend      string `json:"operating_time_weekend"`
}
