package main

import (
	"log"

	"url-shortner/config"
	"url-shortner/controllers"
	"url-shortner/middleware"
	"url-shortner/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	services.InitDB(cfg)
	services.LinkTTLHours = cfg.LinkTTL

	go services.PrePopulateKeys()
	go services.CleanupExpiredLinks()

	r := gin.Default()

	r.Use(middleware.SecurityHeadersMiddleware())
	r.Use(middleware.CORSMiddleware(cfg))

	shortenGroup := r.Group("")
	shortenGroup.Use(middleware.APIKeyAuthMiddleware(cfg))
	shortenGroup.POST("/shorten", controllers.ShortenHandler)

	r.Use(middleware.RateLimitMiddleware(cfg.RateLimit))

	r.GET("/:key", controllers.RedirectHandler)
	r.GET("/links", controllers.ListLinksHandler)

	log.Printf("Server starting on port %s", cfg.Port)
	r.Run(":" + cfg.Port)
}
