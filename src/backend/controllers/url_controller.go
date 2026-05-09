package controllers

import (
	"net/http"
	"net/url"
	"url-shortner/models"
	"url-shortner/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

func ShortenHandler(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	// Basic URL Validation
	_, err := url.ParseRequestURI(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL format"})
		return
	}

	key, err := services.ShortenURL(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to shorten"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"short_url": key})
}

func RedirectHandler(c *gin.Context) {
	key := c.Param("key")
	originalURL, err := services.GetOriginalURL(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func ListLinksHandler(c *gin.Context) {
	collection := services.MongoClient.Database("url_shortner").Collection("urls")
	cursor, err := collection.Find(c, bson.M{}, options.Find().SetLimit(10).SetSort(bson.M{"created_at": -1}))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch links"})
		return
	}
	var results []models.ShortURL
	if err = cursor.All(c, &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode"})
		return
	}
	c.JSON(http.StatusOK, results)
}
