package tests

import (
	"collagen/models/database_model"
	"collagen/tests/test_helper"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAppConfigController(t *testing.T) {
	router := test_helper.SetupServerBeforeTest()

	t.Run("Get App Info", func(t *testing.T) {
		w := test_helper.RequestApiForTest(router, "GET", "/api/v1/app-info", nil)

		// Check if the status code is OK
		assert.Equal(t, http.StatusOK, w.Code)

		// Parse the response body to get the posts
		var appConfig database_model.AppConfigResponse
		err := json.Unmarshal(w.Body.Bytes(), &appConfig)
		if err != nil {
			return
		}
		assert.Equal(t, "", "")
	})

}
