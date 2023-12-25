package post_router

import (
	"collagen/controllers/api/v1/post_controller"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func RegisterPostRouter(v1Router *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	postRouter := v1Router.Group("/post")
	postRouter.Use(authMiddleware.MiddlewareFunc())
	{
		postRouter.GET("/", post_controller.GetPosts)
		postRouter.POST("/", post_controller.CreatePost)
		postRouter.DELETE("/", post_controller.DeletePost)
		postRouter.PUT("/", post_controller.UpdatePost)
	}
}
