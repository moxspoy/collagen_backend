package database_model

import (
	"database/sql"
	"gorm.io/gorm"
)

const (
	StatusVerified = 1
)

type Users struct {
	gorm.Model
	ID                      uint           `json:"id"`
	Email                   sql.NullString `json:"email" gorm:"unique"`
	PhoneNumber             sql.NullString `json:"phone_number" gorm:"unique"`
	Name                    string         `json:"name"`
	Password                string         `json:"password"`
	Pin                     string         `json:"pin"`
	Status                  int            `json:"status"`
	PhoneVerificationStatus int            `json:"phone_verification_status"`
	EmailVerificationStatus int            `json:"email_verification_status"`
}

func (users Users) IsEmailVerified() bool {
	return users.EmailVerificationStatus == StatusVerified
}

func (users Users) IsPhoneVerified() bool {
	return users.PhoneVerificationStatus == StatusVerified
}

func (users Users) IsPinRegistered() bool {
	return users.Pin != ""
}

func (users Users) IsRegistered() bool {
	return users.PhoneNumber.Valid && users.PhoneNumber.String != ""
}
