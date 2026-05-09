package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	MongoURI     string
	ValkeyURL    string
	DBName       string
}

func LoadConfig() *Config {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		Port:      getEnv("PORT", "8080"),
		MongoURI:  getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		ValkeyURL: getEnv("VALKEY_URL", "localhost:6379"),
		DBName:    getEnv("DB_NAME", "url_shortner"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
