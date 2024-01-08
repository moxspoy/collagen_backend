package api_response_model

type SuccessAPIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type SuccessAPIResponseWithToken struct {
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    WithJwtTokenResponse `json:"data"`
}

type SuccessAPIResponseMessageOnly struct {
	Message string `json:"message"`
}

type SuccessAPIResponseMessageAndCode struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
