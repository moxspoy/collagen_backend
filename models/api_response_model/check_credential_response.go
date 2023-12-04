package api_response_model

type CheckCredentialResponse struct {
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"is_email_verified"`
	IsPhoneVerified bool   `json:"is_phone_verified"`
	IsPinRegistered bool   `json:"is_pin_registered"`
	IsRegistered    bool   `json:"is_registered"`
	IsUserExist     bool   `json:"is_user_exist"`
	PhoneNumber     string `json:"phone_number"`
}
