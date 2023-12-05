package middleware

import (
	"database/sql"
	"errors"
	"flop/config/database"
	"flop/models/api_request_model"
	"flop/models/database_model"
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
			if v, ok := data.(*database_model.Users); ok {
				return jwt.MapClaims{
					"email":  v.Email.String,
					"userId": v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			email := claims["email"].(string)
			return &database_model.Users{
				Email: sql.NullString{String: email, Valid: true},
				ID:    uint(claims["userId"].(float64)),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var identityRequest api_request_model.ValidateIdentityRequest
			if err := c.ShouldBind(&identityRequest); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			isPhoneNumber := !(strings.Contains(identityRequest.Credential, "@"))
			var users []database_model.Users
			var whereClause = "email = ?"
			if isPhoneNumber {
				whereClause = "phone_number = ?"
			}
			database.DB.Where(whereClause, identityRequest.Credential).First(&users)

			if len(users) <= 0 {
				return "", errors.New("failed to create token because the user is not found")
			}

			user := users[0]

			if user.ID > 0 {
				return &user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*database_model.Users); ok && v.Name == "admin" {
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

func GetUserIdFromJWT(c *gin.Context) uint {
	claims := jwt.ExtractClaims(c)
	return uint(claims["userId"].(float64))
}

func GetEmailFromJWT(c *gin.Context) string {
	claims := jwt.ExtractClaims(c)
	return claims["email"].(string)
}
