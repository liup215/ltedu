package server

import (
	"edu/conf"
	"edu/lib/net/http/middleware/cors"
	api "edu/server/api" // Keep for now, will refactor handlers
	v1 "edu/server/api/v1"
	"log"

	// "edu/server/pc" // To be removed or merged
	// "edu/server/teacher" // To be removed or merged

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func init() {
	R = gin.Default()

	// Initialize super user before starting server
	if err := v1.EnsureSuperUserExists(); err != nil {
		log.Printf("Failed to initialize super user: %v", err)
		// Don't panic, just log the error
	}

	router(R, conf.Conf)
}

func router(r *gin.Engine, c *conf.Config) {
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("edu_sid", store))
	r.Use(cors.Cors())

	// TODO: The following groups and handlers will be gradually migrated
	// to the unified API or removed if redundant.
	// For now, we keep the backend group as it might contain admin-specific logic
	// that needs careful migration.
	apiGroup := r.Group("/api")    // Path remains /backend for now
	apiHandler := api.NewHandler() // Use the aliased package name
	apiHandler.Route(apiGroup)

}
