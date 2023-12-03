package main

import (
	"flop/config/database"
	v1 "flop/controllers/v1"
	"flop/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	database.ConnectDatabase()

	v1Router := router.Group("/v1")
	v1Router.Use(middleware.GuardApiKey())

	v1Router.GET("/app-info", v1.GetAppInfo)

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

	authRouter.POST("/validate-identity", authMiddleware.LoginHandler)

	// Refresh time can be longer than token timeout
	authRouter.POST("/refresh-token", authMiddleware.RefreshHandler)

	userRouter := v1Router.Group("/user")
	userRouter.Use(authMiddleware.MiddlewareFunc())
	{
		userRouter.GET("/info", v1.GetUserInfo)
		userRouter.PUT("/update-email", v1.UpdateEmail)
	}

	err = router.Run("localhost:8083")
	if err != nil {
		log.Fatal("Error while running server")
	}
}
