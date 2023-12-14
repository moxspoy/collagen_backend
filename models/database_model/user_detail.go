package database_model

import (
	"gorm.io/gorm"
	"time"
)

type UserDetail struct {
	gorm.Model
	UserID                    uint      `form:"user_id"`
	User                      User      `form:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DateOfBirth               time.Time `form:"date_of_birth" gorm:"default:1900-01-01 01:01:01.111"`
	Gender                    string    `form:"gender"`
	AddressLine1              string    `form:"address_line_1"`
	AddressLine2              string    `form:"address_line_2"`
	City                      string    `form:"city"`
	State                     string    `form:"state"`
	Country                   string    `form:"country"`
	ZipCode                   int       `form:"zip_code"`
	DocumentType              string    `form:"document_type"`
	DocumentNumber            string    `form:"document_number"`
	SelfieImageURL            string    `form:"selfie_image_url"`
	IdentityImageURL          string    `form:"identity_image_url"`
	IdentityAndSelfieImageURL string    `form:"identity_and_selfie_image_url"`
	IssueDate                 time.Time `form:"issue_date"  gorm:"default:1900-01-01 01:01:01.111"`
	ExpiryDate                time.Time `form:"expiry_date"  gorm:"default:1900-01-01 01:01:01.111"`
}

func (UserDetail) TableName() string {
	return "user_detail"
}