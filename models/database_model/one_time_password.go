package database_model

import (
	"gorm.io/gorm"
	"time"
)

type OneTimePassword struct {
	gorm.Model
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	Status    int       `json:"status"`
}

func (OneTimePassword) TableName() string {
	return "one_time_password"
}

func (otp OneTimePassword) IsOTPExpired() bool {
	return otp.Status == 2
}
