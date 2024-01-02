package test_helper

import (
	"collagen/config/database"
	"collagen/routes/central_router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

func SetupServerBeforeTest() *gin.Engine {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize a test router
	router := gin.Default()

	// Set the test mode to disable unnecessary output
	gin.SetMode(gin.TestMode)
	central_router.RegisterAllRouter(router)

	// Create a temporary in-memory database for testing
	database.DB, _ = database.SetupTestDB()
	err = database.RunAutoMigration(database.DB)
	if err != nil {
		return nil
	}
	return router
}

func RequestApiForTest(router *gin.Engine, method string, url string, body io.Reader) *httptest.ResponseRecorder {
	requiredApiKey := os.Getenv("API_KEY")

	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_key", requiredApiKey)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
