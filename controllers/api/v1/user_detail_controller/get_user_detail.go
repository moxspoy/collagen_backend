package user_detail_controller

import (
	"flop/middleware"
	"flop/models/database_model"
	"flop/repositories/user_detail_repository"
	"flop/repositories/user_repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserDetail godoc
//
//	@Summary		Get user detail object
//	@Description	This endpoint used to fetch user's data but with more detail
//	@Tags			User
//	@Produce		json
//	@Success		200	{object} database_model.User
//	@Router			/user-detail/info [get]
//	@Param			api_key	header string	true "Api Key"
//	@Security		ApiKeyAuth
func GetUserDetail(c *gin.Context) {
	userId := middleware.GetUserIdFromJWT(c)

	currentUser := database_model.UserDetail{
		UserID: userId,
	}
	if currentUser.ID <= 0 {
		user := user_repository.GetOneUserById(userId)
		currentUser.User = user
	}
	user_detail_repository.GetUserDetail(&currentUser)
	c.IndentedJSON(http.StatusOK, currentUser)
}
