package main

import (
	"flop/config/database"
	"flop/routes/central_router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

// @title           Flop Web Service
// @version         1.0
// @description     Web service API in Go using Gin framework.
// @termsOfService  https://tos.santoshk.dev

// @contact.name   M Nurilman Baehaqi
// @contact.url    https://twitter.com/MOXSPOY
// @contact.email  mnurilmanbaehaqi@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8083
// @BasePath  /api/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.SetupDatabase()

	router := gin.Default()
	central_router.RegisterAllRouter(router)

	err = router.Run("localhost:8083")
	if err != nil {
		log.Fatal("Error while running server")
	}
}
