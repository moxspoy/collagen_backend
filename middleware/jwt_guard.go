package middleware

import (
	"flop/config/database"
	"flop/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
)

const JwtIdentityKey = "id"

func GetJWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(os.Getenv("JWT_KEY")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: JwtIdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Users); ok {
				return jwt.MapClaims{
					JwtIdentityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.Users{
				Email: claims[JwtIdentityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var identityRequest models.IdentityRequest
			if err := c.ShouldBind(&identityRequest); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			isPhoneNumber := !(strings.Contains(identityRequest.Credential, "@"))
			var users []models.Users
			var whereClause = "email = ?"
			if isPhoneNumber {
				whereClause = "phone_number = ?"
			}
			database.DB.Where(whereClause, identityRequest.Credential).First(&users)

			if len(users) <= 0 {
				return "", jwt.ErrFailedTokenCreation
			}

			user := users[0]

			if user.Id > 0 {
				return &user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.Users); ok && v.Name == "admin" {
				return true
			}

			return true
			//return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.IndentedJSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}
