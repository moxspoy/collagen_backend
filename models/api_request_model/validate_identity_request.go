package api_request_model

type ValidateIdentityRequest struct {
	Credential       string `form:"credential"`
	DeviceIdentifier string `form:"device_identifier"`
	DeviceModel      string `form:"device_model"`
	OsVersion        string `form:"os_version"`
	Platform         string `form:"platform"`
	RequestId        string `form:"request_id"`
	AppNameVersion   string `form:"app_name_version"`
	AppBuildVersion  int    `form:"app_build_version"`
}
