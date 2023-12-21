package universities_router

import (
	"collagen/controllers/api/v1/universities_controller"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func RegisterUniversitiesRouter(v1Router *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	universitiesRouter := v1Router.Group("/universities")
	universitiesRouter.Use(authMiddleware.MiddlewareFunc())
	{
		universitiesRouter.GET("/all", universities_controller.GetAllUniversities)
		universitiesRouter.POST("/create", universities_controller.CreateUniversity)
	}
}
