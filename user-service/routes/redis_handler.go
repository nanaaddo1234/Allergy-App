package routes

import (
	"log"
	"net/http"
	"time"
	"user-service/utils" // make sure this matches your actual module name

	"github.com/gin-gonic/gin"
)

func RedisHandler(c *gin.Context) {
	// Extract request-scoped context
	ctx := c.Request.Context()

	// Write to Redis with cancellation support
	err := utils.RedisClient.Set(ctx, "key", "value", time.Hour).Err()
	if err != nil {
		log.Println("Redis failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
