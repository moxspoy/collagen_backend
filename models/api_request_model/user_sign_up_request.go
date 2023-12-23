package api_request_model

type UserSignUpRequest struct {
	Name             string `form:"name"`
	Email            string `form:"email"`
	PhoneNumber      string `form:"phoneNumber"`
	Credential       string `form:"credential"`
	Platform         string `form:"platform"`
	DeviceModel      string `form:"deviceModel"`
	AppNameVersion   string `form:"appNameVersion"`
	AppBuildVersion  int    `form:"appBuildVersion"`
	OsVersion        string `form:"osVersion"`
	DeviceIdentifier string `form:"deviceIdentifier"`
}
