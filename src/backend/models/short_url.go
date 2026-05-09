package models

import "time"

type ShortURL struct {
	ID        string    `bson:"_id"`
	Original  string    `bson:"original"`
	CreatedAt time.Time `bson:"created_at"`
	ExpiresAt time.Time `bson:"expires_at"`
}
