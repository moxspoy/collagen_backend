package central_router

import (
	"collagen/controllers/api/v1/auth_controller"
	"collagen/middleware"
	"collagen/routes/auth_router"
	"collagen/routes/otp_router"
	"collagen/routes/post_router"
	"collagen/routes/public_router"
	"collagen/routes/swagger_router"
	"collagen/routes/universities_router"
	"collagen/routes/user_detail_router"
	"collagen/routes/user_router"
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
	universities_router.RegisterUniversitiesRouter(v1Router, authMiddleware)
	post_router.RegisterPostRouter(v1Router, authMiddleware)
	swagger_router.RegisterSwaggerRouter(router)
}
