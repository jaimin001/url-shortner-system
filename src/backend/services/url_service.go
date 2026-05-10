package services

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"
	"url-shortner/models"

	"go.mongodb.org/mongo-driver/bson"
)

var LinkTTLHours int = 24

func PrePopulateKeys() {
	for {
		size, err := ValkeyClient.LLen(context.TODO(), "available_keys").Result()
		if err != nil {
			log.Println("Error checking key list size:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if size < 1000 {
			for i := 0; i < 100; i++ {
				key := generateKey(6)
				ValkeyClient.RPush(context.TODO(), "available_keys", key)
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func ShortenURL(originalURL string) (string, error) {
	var key string
	key, err := ValkeyClient.LPop(context.TODO(), "available_keys").Result()
	if err != nil {
		key = generateKey(6)
	}

	ttl := time.Duration(LinkTTLHours) * time.Hour
	expiresAt := time.Now().Add(ttl)

	newURL := models.ShortURL{
		ID:        key,
		Original:  originalURL,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
	}

	collection := MongoClient.Database("url_shortner").Collection("urls")
	_, err = collection.InsertOne(context.TODO(), newURL)
	if err != nil {
		return "", err
	}

	err = ValkeyClient.Set(context.TODO(), key, originalURL, ttl).Err()
	if err != nil {
		return key, nil
	}

	return key, nil
}

func GetOriginalURL(key string) (string, error) {
	val, err := ValkeyClient.Get(context.TODO(), key).Result()
	if err == nil {
		return val, nil
	}

	collection := MongoClient.Database("url_shortner").Collection("urls")
	var result models.ShortURL
	err = collection.FindOne(context.TODO(), bson.M{"_id": key}).Decode(&result)
	if err != nil {
		return "", err
	}

	if time.Now().After(result.ExpiresAt) {
		return "", errors.New("link has expired")
	}

	ttl := time.Until(result.ExpiresAt)
	if ttl > 0 {
		ValkeyClient.Set(context.TODO(), key, result.Original, ttl)
	}

	return result.Original, nil
}

func CleanupExpiredLinks() {
	for {
		collection := MongoClient.Database("url_shortner").Collection("urls")
		filter := bson.M{"expires_at": bson.M{"$lt": time.Now()}}

		_, err := collection.DeleteMany(context.TODO(), filter)
		if err != nil {
			log.Println("Error cleaning up expired links:", err)
		}

		time.Sleep(1 * time.Hour)
	}
}

func generateKey(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
