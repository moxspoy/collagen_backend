package api_request_model

type ValidateIdentityRequest struct {
	Credential       string `form:"credential"`
	DeviceIdentifier string `form:"deviceIdentifier"`
	DeviceModel      string `form:"deviceModel"`
	OsVersion        string `form:"osVersion"`
	Platform         string `form:"platform"`
	RequestId        string `form:"requestId"`
	AppNameVersion   string `form:"appNameVersion"`
	AppBuildVersion  int    `form:"appBuildVersion"`
}
