package main

import (
	"log"
	"url-shortner/config"
	"url-shortner/controllers"
	"url-shortner/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize infrastructure
	services.InitDB(cfg)

	// Start background worker for pre-populating keys
	go services.PrePopulateKeys()

	// Setup Router
	r := gin.Default()

	// Add CORS middleware first
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/shorten", controllers.ShortenHandler)
	r.GET("/:key", controllers.RedirectHandler)
	r.GET("/links", controllers.ListLinksHandler)


	log.Printf("Server starting on port %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
