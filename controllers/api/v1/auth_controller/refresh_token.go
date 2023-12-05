package auth_controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// RefreshToken godoc
//
//	@Summary		Refresh token
//	@Description	Use this endpoint to refresh token
//	@Tags			Auth
//	@Produce		json
//	@Success		200	{object} api_response_model.JwtResponse
//	@Router			/auth/refresh-token [post]
//	@Param			api_key	header string	true "Api Key"
//	@Security		ApiKeyAuth
func RefreshToken(c *gin.Context, authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware.RefreshHandler(c)
}
