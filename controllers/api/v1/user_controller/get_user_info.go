package user_controller

import (
	"flop/middleware"
	"flop/repositories/user_repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserInfo godoc
//
//	@Summary		Get user object
//	@Description	This endpoint used to fetch user's data
//	@Tags			User
//	@Produce		json
//	@Success		200	{object} database_model.User
//	@Router			/user/info [get]
//	@Param			api_key	header string	true "Api Key"
//	@Security		ApiKeyAuth
func GetUserInfo(c *gin.Context) {
	userId := middleware.GetUserIdFromJWT(c)
	currentUser := user_repository.GetOneUserById(userId)
	c.IndentedJSON(http.StatusOK, currentUser)
}
