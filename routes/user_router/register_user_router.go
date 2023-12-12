package user_router

import (
	"flop/controllers/api/v1/user_controller"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(v1Router *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	userRouter := v1Router.Group("/user")
	userRouter.Use(authMiddleware.MiddlewareFunc())
	{
		userRouter.GET("/info", user_controller.GetUserInfo)
		userRouter.PUT("/update-phone-number", user_controller.UpdatePhoneNumber)
		userRouter.PUT("/update-user-name", user_controller.UpdateUserName)
		userRouter.PUT("/update-pin", user_controller.UpdatePin)
		userRouter.PUT("/update-email", func(context *gin.Context) {
			user_controller.UpdateEmail(context, authMiddleware)
		})
	}
}
