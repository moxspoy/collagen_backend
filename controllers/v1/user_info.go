package v1

import (
	"flop/middleware"
	"flop/repositories/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserInfo godoc
//
//	@Summary		Get user object
//	@Description	This endpoint used to fetch user's data
//	@Tags			User
//	@Produce		json
//	@Success		200	{object} database_model.Users
//	@Router			/user/info [get]
//	@Param			api_key	header string	true "Api Key"
//	@Security		ApiKeyAuth
func GetUserInfo(c *gin.Context) {
	userId := middleware.GetUserIdFromJWT(c)
	currentUser := users.GetOneUserById(userId)
	c.IndentedJSON(http.StatusOK, currentUser)
}
