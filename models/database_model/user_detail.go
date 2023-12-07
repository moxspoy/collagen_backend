package database_model

import (
	"gorm.io/gorm"
	"time"
)

type UserDetail struct {
	gorm.Model
	UserID                    uint      `json:"user_id"`
	User                      Users     `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	DateOfBirth               time.Time `json:"date_of_birth"`
	Gender                    string    `json:"gender" gorm:"check:gender IN ('Male', 'Female', 'Other')"`
	AddressLine1              string    `json:"address_line_1"`
	AddressLine2              string    `json:"address_line_2"`
	City                      string    `json:"city"`
	State                     string    `json:"state"`
	Country                   string    `json:"country"`
	ZipCode                   int       `json:"zip_code"`
	DocumentType              string    `json:"document_type" gorm:"check:document_type IN ('Passport', 'Driver''s License', 'ID Card')"`
	DocumentNumber            string    `json:"document_number"`
	SelfieImageURL            string    `json:"selfie_image_url"`
	IdentityImageURL          string    `json:"identity_image_url"`
	IdentityAndSelfieImageURL string    `json:"identity_and_selfie_image_url"`
	IssueDate                 time.Time `json:"issue_date"`
	ExpiryDate                time.Time `json:"expiry_date"`
}

func (UserDetail) TableName() string {
	return "user_detail"
}
