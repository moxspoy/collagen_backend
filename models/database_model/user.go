package database_model

import (
	"database/sql"
	"gorm.io/gorm"
)

const (
	StatusVerified = 1
)

type User struct {
	gorm.Model
	Email                   sql.NullString `json:"email" gorm:"unique"`
	PhoneNumber             sql.NullString `json:"phone_number" gorm:"unique"`
	Name                    string         `json:"name"`
	Password                string         `json:"-"`
	Pin                     string         `json:"-"`
	Status                  int            `json:"status"`
	PhoneVerificationStatus int            `json:"phone_verification_status"`
	EmailVerificationStatus int            `json:"email_verification_status"`
}

func (User) TableName() string {
	return "user"
}

func (user User) IsEmailVerified() bool {
	return user.EmailVerificationStatus == StatusVerified
}

func (user User) IsPhoneVerified() bool {
	return user.PhoneVerificationStatus == StatusVerified
}

func (user User) IsPinRegistered() bool {
	return user.Pin != ""
}

func (user User) IsRegistered() bool {
	return user.PhoneNumber.Valid && user.PhoneNumber.String != ""
}
