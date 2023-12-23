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
	Email                   sql.NullString `gorm:"unique"`
	PhoneNumber             sql.NullString `gorm:"unique"`
	Name                    string
	Password                string `json:"-"`
	Pin                     string `json:"-"`
	Status                  int
	PhoneVerificationStatus int
	EmailVerificationStatus int
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
