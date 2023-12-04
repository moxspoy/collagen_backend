package models

const (
	StatusVerified = 1
)

type Users struct {
	Id                      uint   `json:"id"`
	Email                   string `json:"email"`
	PhoneNumber             string `json:"phone_number"`
	Name                    string `json:"name"`
	Password                string `json:"password"`
	Pin                     string `json:"pin"`
	Status                  int    `json:"status"`
	PhoneVerificationStatus int    `json:"phone_verification_status"`
	EmailVerificationStatus int    `json:"email_verification_status"`
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
