package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	MongoURI      string
	ValkeyURL     string
	DBName        string
	AllowedOrigin string
	RateLimit     int
	APIKey        string
	LinkTTL       int
}

func LoadConfig() *Config {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		Port:          getEnv("PORT", "8080"),
		MongoURI:      getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		ValkeyURL:     getEnv("VALKEY_URL", "localhost:6379"),
		DBName:        getEnv("DB_NAME", "url_shortner"),
		AllowedOrigin: getEnv("ALLOWED_ORIGIN", "http://localhost:5173"),
		RateLimit:     getEnvInt("RATE_LIMIT", 30),
		APIKey:        getEnv("API_KEY", ""),
		LinkTTL:       getEnvInt("LINK_TTL", 24),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return fallback
}
