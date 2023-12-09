package auth_router

import (
	"flop/controllers/api/v1/auth_controller"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRouter(authRouter *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	authRouter.POST("/validate-identity", func(context *gin.Context) {
		auth_controller.ValidateIdentity(context, authMiddleware)
	})
	authRouter.POST("/sign-up", func(context *gin.Context) {
		auth_controller.UserSignUp(context, authMiddleware)
	})
	authRouter.POST("/refresh-token", func(context *gin.Context) {
		auth_controller.RefreshToken(context, authMiddleware)
	})

}
