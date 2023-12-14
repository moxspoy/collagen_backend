package otp_router

import (
	"collagen/controllers/api/v1/otp_controller"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func RegisterOtpRouter(v1Router *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	otpRouter := v1Router.Group("/otp")
	otpRouter.Use(authMiddleware.MiddlewareFunc())
	{
		otpRouter.POST("/request", otp_controller.RequestOTP)
		otpRouter.POST("/check", otp_controller.CheckOTP)
	}
}
