package models

type SuccessAPIResponse struct {
	code    int32
	message string
	data    any
}

type SuccessAPIResponseMessageOnly struct {
	Message string `json:"message"`
}

type SuccessAPIResponseMessageAndCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
