package cors

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func Cors() gin.HandlerFunc {
	return cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders: "Origin, Authorization, Content-Type, token, x_requested_with",
		ExposedHeaders: "",
		MaxAge:         50 * time.Second,
		// Credentials:     true,
		ValidateHeaders: false,
	})
}
