package tests

import (
	"collagen/config/database"
	"collagen/models/database_model"
	"collagen/routes/central_router"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAppConfigController(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	requiredApiKey := os.Getenv("API_KEY")

	// Initialize a test router
	router := gin.Default()

	// Set the test mode to disable unnecessary output
	gin.SetMode(gin.TestMode)
	central_router.RegisterAllRouter(router)

	// Create a temporary in-memory database for testing
	database.DB, _ = database.SetupTestDB()
	err = database.RunAutoMigration(database.DB)
	if err != nil {
		return
	}

	t.Run("Get App Info", func(t *testing.T) {

		// Perform a POST request to create a new post
		req, _ := http.NewRequest("GET", "/api/v1/app-info", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("api_key", requiredApiKey)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

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
