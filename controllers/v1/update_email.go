package v1

import (
	"flop/helper/api_response"
	"flop/repositories/users"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func UpdateEmail(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	oldEmail := fmt.Sprintf("%v", claims[identityKey])
	newEmail := c.Request.FormValue("new_email")
	users.UpdateUserEmail(oldEmail, newEmail)
	api_response.GenerateSuccessResponse(c, "Update email successful")
}
