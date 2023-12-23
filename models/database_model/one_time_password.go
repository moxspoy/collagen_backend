package database_model

import (
	"gorm.io/gorm"
	"time"
)

type OneTimePassword struct {
	gorm.Model
	UserID    uint
	User      User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Token     string `json:"-"`
	ExpiredAt time.Time
	Status    int
}

func (OneTimePassword) TableName() string {
	return "one_time_password"
}

func (otp OneTimePassword) IsOTPExpired() bool {
	return otp.Status == 2
}
