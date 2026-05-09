package services

import (
	"context"
	"log"
	"math/rand"
	"time"
	"url-shortner/models"

	"go.mongodb.org/mongo-driver/bson"
)

func PrePopulateKeys() {
	for {
		// Check current size of the list
		size, err := ValkeyClient.LLen(context.TODO(), "available_keys").Result()
		if err != nil {
			log.Println("Error checking key list size:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Maintain a minimum buffer of 1000 keys
		if size < 1000 {
			for i := 0; i < 100; i++ {
				key := generateKey(6)
				ValkeyClient.RPush(context.TODO(), "available_keys", key)
			}
		}
		time.Sleep(1 * time.Second) // Avoid tight loop
	}
}

func ShortenURL(originalURL string) (string, error) {
	var key string
	// Try to pop a key
	key, err := ValkeyClient.LPop(context.TODO(), "available_keys").Result()
	if err != nil {
		// Fallback to random generation if list is empty or redis fails
		key = generateKey(6)
	}

	newURL := models.ShortURL{
		ID:        key,
		Original:  originalURL,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	// Save to MongoDB
	collection := MongoClient.Database("url_shortner").Collection("urls")
	_, err = collection.InsertOne(context.TODO(), newURL)
	if err != nil {
		return "", err
	}

	// Save to Valkey
	err = ValkeyClient.Set(context.TODO(), key, originalURL, 24*time.Hour).Err()
	if err != nil {
		// Just log error if cache fails, as Mongo is the source of truth
		return key, nil
	}

	return key, nil
}

func GetOriginalURL(key string) (string, error) {
	// Try Valkey first
	val, err := ValkeyClient.Get(context.TODO(), key).Result()
	if err == nil {
		return val, nil
	}

	// If not in Valkey, check MongoDB
	collection := MongoClient.Database("url_shortner").Collection("urls")
	var result models.ShortURL
	err = collection.FindOne(context.TODO(), bson.M{"_id": key}).Decode(&result)
	if err != nil {
		return "", err
	}

	// Backfill Valkey
	ValkeyClient.Set(context.TODO(), key, result.Original, 24*time.Hour)

	return result.Original, nil
}

func generateKey(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
