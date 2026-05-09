package services

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/redis/go-redis/v9"
	"url-shortner/config"
)

var (
	MongoClient *mongo.Client
	ValkeyClient *redis.Client
)

func InitDB(cfg *config.Config) {
	// Initialize MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}
	MongoClient = client

	// Initialize Valkey
	ValkeyClient = redis.NewClient(&redis.Options{
		Addr: cfg.ValkeyURL,
	})

	log.Println("Successfully connected to Database and Valkey")
}
