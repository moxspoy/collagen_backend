package models

type IdentityRequest struct {
	Credential       string `form:"credential"`
	DeviceIdentifier string `form:"deviceIdentifier"`
	DeviceModel      string `form:"deviceModel"`
	OsVersion        int    `form:"osVersion"`
	Platform         string `form:"platform"`
	RequestId        string `form:"requestId"`
	Version          int    `form:"version"`
}
