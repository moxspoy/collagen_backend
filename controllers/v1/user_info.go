package v1

import (
	"flop/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

func GetUserInfo(c *gin.Context) {

	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.IndentedJSON(200, gin.H{
		"email":    claims[identityKey],
		"userName": user.(*models.Users).Name,
		"text":     "Hello World.",
	})

	//isPhoneNumber := !(strings.Contains(claims[identityKey], "@"))
	//var users []models.Users
	//var whereClause = "email = ?"
	//if isPhoneNumber {
	//	whereClause = "phone_number = ?"
	//}
	//database.DB.Where(whereClause, identityRequest.Credential).First(&users)
	//
	//if len(users) <= 0 {
	//	return "", jwt.ErrFailedTokenCreation
	//}
	//
	//user := users[0]
}
