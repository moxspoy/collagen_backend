package test_helper

import (
	"collagen/config/database"
	"collagen/models/database_model"
	"collagen/routes/central_router"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
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

func RequestPublicApiForTest(router *gin.Engine, method string, url string, body io.Reader) *httptest.ResponseRecorder {
	requiredApiKey := os.Getenv("API_KEY")

	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_key", requiredApiKey)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func RequestApiForTest(router *gin.Engine, method string, url string, body io.Reader) *httptest.ResponseRecorder {
	requiredApiKey := os.Getenv("API_KEY")

	// Simulate user authentication and generate a valid JWT token
	user := &database_model.User{
		Model: gorm.Model{
			ID: 1,
		},
		Email: sql.NullString{
			String: "ikrimah@gmail.com",
			Valid:  true,
		},
	}

	// Generate a valid JWT token for the user
	identityKey := "id"
	signingKey := []byte(os.Getenv("JWT_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		identityKey: user.ID,
		"userId":    user.ID,
		"email":     user.Email.String,
	})
	tokenString, _ := token.SignedString(signingKey)

	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api_key", requiredApiKey)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
