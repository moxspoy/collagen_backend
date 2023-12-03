package v1

import (
	"flop/repositories/users"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var identityKey = "id"

// GetUserInfo godoc
//
//	@Summary		Get user object
//	@Description	This endpoint used to fetch user's data
//	@Tags			User
//	@Produce		json
//	@Success		200	{object} models.Users
//	@Router			/user/info [get]
//	@Param			api_key	header string	true "Api Key"
//	@Security		ApiKeyAuth
func GetUserInfo(c *gin.Context) {

	claims := jwt.ExtractClaims(c)
	//user, _ := c.Get(identityKey)
	email := fmt.Sprintf("%v", claims[identityKey])
	currentUser := users.GetOneUserByEmail(email)
	log.Print(currentUser)
	c.IndentedJSON(http.StatusOK, currentUser)
}
