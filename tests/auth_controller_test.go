package tests

import (
	"collagen/models/api_request_model"
	"collagen/models/api_response_model"
	"collagen/tests/test_helper"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestAuthController(t *testing.T) {
	router := test_helper.SetupServerBeforeTest()
	var token string

	t.Run("Sign Up", func(t *testing.T) {
		request := api_request_model.UserSignUpRequest{
			Name:             "Eqi",
			Email:            "eqi@gmail.com",
			PhoneNumber:      "",
			Credential:       "eqi@gmail.com",
			Platform:         "android",
			DeviceModel:      "Mi 10T Pro",
			AppNameVersion:   "1.0.0",
			AppBuildVersion:  1,
			OsVersion:        "12",
			DeviceIdentifier: "DeviceIdentifier",
		}
		requestJSON, _ := json.Marshal(request)

		w := test_helper.RequestApiWithoutTokenForTest(router, "POST", "/api/v1/auth/sign-up", strings.NewReader(string(requestJSON)))

		assert.Equal(t, http.StatusOK, w.Code)

		var response api_response_model.SuccessAPIResponseWithToken
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			return
		}

		token = response.Data.Token
		assert.Equal(t, response.Message, "success register")
	})

	t.Run("Refresh Token", func(t *testing.T) {

		w := test_helper.RequestApiWithTokenForTest(router, "POST", "/api/v1/auth/refresh-token", nil, token)

		var response api_response_model.JwtResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			return
		}

		assert.Equal(t, http.StatusOK, w.Code)
		newToken := response.Token
		fmt.Println("registr_token: ", token)
		fmt.Println("refresh_token: ", newToken)
	})

	t.Run("Sign In", func(t *testing.T) {

		request := api_request_model.ValidateIdentityRequest{
			Credential:       "eqi@gmail.com",
			Platform:         "android",
			DeviceModel:      "Mi 10T Pro",
			AppNameVersion:   "1.0.0",
			AppBuildVersion:  1,
			OsVersion:        "12",
			DeviceIdentifier: "DeviceIdentifier",
			RequestId:        "RequestId",
		}
		requestJSON, _ := json.Marshal(request)

		w := test_helper.RequestApiWithoutTokenForTest(router, "POST", "/api/v1/auth/validate-identity", strings.NewReader(string(requestJSON)))

		var response api_response_model.JwtResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		if err != nil {
			return
		}

		assert.Equal(t, http.StatusOK, w.Code)
		if response.Token == "" {
			t.Error("Token is empty")
		}
		newToken := response.Token
		fmt.Println("signin_token: ", newToken)
	})

}
