package main

import (
	"flop/config/database"
	v1 "flop/controllers/v1"
	"flop/controllers/v1/public_controller"
	"flop/docs"
	"flop/middleware"
	"flop/models/database_model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router := gin.Default()
	database.ConnectDatabase()
	err = database.DB.AutoMigrate(&database_model.Users{}, &database_model.AppConfig{}, &database_model.UserLoggedInDevices{})
	if err != nil {
		log.Fatal("Migration Error:" + err.Error())
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	apiRouter := router.Group("/api")
	v1Router := apiRouter.Group("/v1")
	v1Router.Use(middleware.GuardApiKey())

	v1Router.GET("/app-info", public_controller.GetAppInfo)

	authRouter := v1Router.Group("auth")
	authRouter.POST("/check-credential", v1.CheckCredential)

	authMiddleware, err := middleware.GetJWTMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	authRouter.POST("/validate-identity", func(context *gin.Context) {
		v1.ValidateIdentity(context, authMiddleware)
	})
	authRouter.POST("/sign-up", func(context *gin.Context) {
		v1.UserSignUp(context, authMiddleware)
	})

	// Refresh time can be longer than token timeout
	authRouter.POST("/refresh-token", func(context *gin.Context) {
		v1.RefreshToken(context, authMiddleware)
	})

	userRouter := v1Router.Group("/user")
	userRouter.Use(authMiddleware.MiddlewareFunc())
	{
		userRouter.GET("/info", v1.GetUserInfo)
		userRouter.PUT("/update-email", func(context *gin.Context) {
			v1.UpdateEmail(context, authMiddleware)
		})
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err = router.Run("localhost:8083")
	if err != nil {
		log.Fatal("Error while running server")
	}
}
