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
