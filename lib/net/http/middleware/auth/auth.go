package auth

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func New(c *Config) (*jwt.GinJWTMiddleware, error) {
	if c == nil {
		c = &Config{}
	}

	if c.Realm == "" {
		c.Realm = "ltedu"
	}

	if c.Key == "" {
		c.Key = "ltedu"
	}

	if c.TokenLookup == "" {
		c.TokenLookup = "header: Authorization, query: token, cookie: jwt"
	}

	if c.TokenHeadName == "" {
		c.TokenHeadName = "Bearer"
	}

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:   c.Realm,
		Key:     []byte(c.Key),
		Timeout: 24 * 30 * time.Hour,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		// TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenLookup: c.TokenLookup,
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		// TokenHeadName: "Bearer",
		TokenHeadName: c.TokenHeadName,

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	})

	return authMiddleware, err
}
