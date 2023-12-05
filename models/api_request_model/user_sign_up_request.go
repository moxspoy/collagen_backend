package api_request_model

type UserSignUpRequest struct {
	Name             string `form:"name"`
	Email            string `form:"email"`
	PhoneNumber      string `form:"phone_number"`
	Credential       string `form:"credential"`
	Platform         string `form:"platform"`
	DeviceModel      string `form:"device_model"`
	AppNameVersion   string `form:"app_name_version"`
	AppBuildVersion  int    `form:"app_build_version"`
	OsVersion        string `form:"os_version"`
	DeviceIdentifier string `form:"device_identifier"`
}
