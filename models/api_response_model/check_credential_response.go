package api_response_model

type CheckCredentialResponse struct {
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"isEmailVerified"`
	IsPhoneVerified bool   `json:"isPhoneVerified"`
	IsPinRegistered bool   `json:"isPinRegistered"`
	IsRegistered    bool   `json:"isRegistered"`
	IsUserExist     bool   `json:"isUserExist"`
	PhoneNumber     string `json:"phoneNumber"`
}
