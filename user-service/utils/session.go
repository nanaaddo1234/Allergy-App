package utils

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var RedisClient *redis.Client

// InitRedis initializes a connection to the Redis server.
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // change if using Docker or a different host
		Password: "",               // or your Redis password
		DB:       0,
	})

	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

// CreateRefreshToken generates a UUID-based refresh token and stores it in Redis with a 7-day expiration.
func CreateRefreshToken(userID uint) (string, error) {
	token := uuid.New().String()
	err := RedisClient.Set(ctx, token, userID, 7*24*time.Hour).Err() // 7 days
	return token, err
}

// GetUserIDFromToken retrieves the user ID associated with a given refresh token.
func GetUserIDFromToken(token string) (string, error) {
	return RedisClient.Get(ctx, token).Result()
}

// RotateRefreshToken deletes the old refresh token and creates a new one for the given user ID.
func RotateRefreshToken(oldToken string, userID uint) (string, error) {
	RedisClient.Del(ctx, oldToken)
	return CreateRefreshToken(userID)
}

// DeleteRefreshToken deletes a refresh token from Redis, typically during logout.
func DeleteRefreshToken(token string) {
	if err := RedisClient.Del(ctx, token).Err(); err != nil {
		log.Printf("Error deleting refresh token: %v", err)
	}
}
