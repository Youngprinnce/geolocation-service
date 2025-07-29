package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"github.com/youngprinnce/geolocation-service/config"
	"github.com/youngprinnce/geolocation-service/internal/app"
	"github.com/youngprinnce/geolocation-service/internal/app/manualwire"
	"github.com/youngprinnce/geolocation-service/internal/logger"
)

func init() {
	log.SetLevel(log.WarnLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func RegisterRoutes(conf *config.Config) *gin.Engine {
	binding.Validator = new(app.DefaultValidator)

	router := gin.Default()
	router.MaxMultipartMemory = 2 << 20 // 2 MiB

	// CORS middleware
	router.Use(corsMiddleware())
	router.Use(gin.Recovery())
	router.Use(errorReporterMiddleware())

	// Health and info endpoints
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello!")
	})

	locationController := manualwire.GetLocationController()

	// Location routes
	locationRoutes := router.Group("/locations")
	{
		locationRoutes.POST("", locationController.CreateLocation)
		locationRoutes.GET("", locationController.GetLocations)
		locationRoutes.GET("/nearest", locationController.GetNearest)
		locationRoutes.DELETE("/:name", locationController.DeleteLocation)
	}

	logger.Info("App routes registered successfully!")

	return router
}

// corsMiddleware handles CORS headers
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func errorReporterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			log.WithFields(log.Fields{
				"method": c.Request.Method,
				"path":   c.Request.URL.Path,
				"errors": c.Errors.String(),
			}).Error("Request completed with errors")
		}
	}
}
