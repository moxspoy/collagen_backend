package auth_controller

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// ValidateIdentity godoc
//
//	@Summary		Validate identity is like for login
//	@Description	This endpoint used to check user in the platform (for login)
//	@Tags			Auth
//	@Produce		json
//	@Accept			multipart/form-data
//	@Success		200	{object} api_response_model.JwtResponse
//	@Router			/auth/validate-identity [post]
//	@Param			api_key	header string	true "Api Key"
//	@Param			data formData api_request_model.ValidateIdentityRequest true "Request object"
func ValidateIdentity(c *gin.Context, authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware.LoginHandler(c)
}
