package main

import (
	"flop/config/database"
	v1 "flop/controllers/v1"
	"flop/middleware"
	"flop/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"strings"
	"time"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*models.Users).Name,
		"text":     "Hello World.",
	})
}

var identityKey = "id"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	database.ConnectDatabase()

	v1Router := router.Group("/v1")
	v1Router.Use(middleware.GuardApiKey())

	v1Router.GET("/app-info", v1.GetAppInfo)

	authRouter := v1Router.Group("auth")
	authRouter.POST("/check-credential", v1.CheckCredential)

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secretx_xxkey"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.Users); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.Users{
				Email: claims[identityKey].(string),
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

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	authRouter.POST("/validate-identity", authMiddleware.LoginHandler)

	// Refresh time can be longer than token timeout
	authRouter.POST("/refresh-token", authMiddleware.RefreshHandler)

	userRouter := v1Router.Group("/user")
	userRouter.Use(authMiddleware.MiddlewareFunc())
	{
		userRouter.GET("/info", helloHandler)
	}

	err = router.Run("localhost:8083")
	if err != nil {
		log.Fatal("Error while running server")
	}
}
