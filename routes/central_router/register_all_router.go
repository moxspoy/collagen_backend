package central_router

import (
	"flop/controllers/api/v1/auth_controller"
	"flop/middleware"
	"flop/routes/auth_router"
	"flop/routes/otp_router"
	"flop/routes/public_router"
	"flop/routes/swagger_router"
	"flop/routes/user_detail_router"
	"flop/routes/user_router"
	"github.com/gin-gonic/gin"
	"log"
)

func RegisterAllRouter(router *gin.Engine) {
	apiRouter := router.Group("/api")

	v1Router := apiRouter.Group("/v1")
	public_router.RegisterPublicRouter(v1Router)

	v1Router.Use(middleware.GuardApiKey())

	// Auth
	authRouter := v1Router.Group("auth")
	authRouter.POST("/check-credential", auth_controller.CheckCredential)

	authMiddleware, err := middleware.GetJWTMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errMiddleware := authMiddleware.MiddlewareInit()
	if errMiddleware != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errMiddleware.Error())
	}
	auth_router.RegisterAuthRouter(v1Router, authMiddleware)

	user_router.RegisterUserRouter(v1Router, authMiddleware)
	user_detail_router.RegisterUserDetailRouter(v1Router, authMiddleware)
	otp_router.RegisterOtpRouter(v1Router, authMiddleware)
	swagger_router.RegisterSwaggerRouter(router)
}
