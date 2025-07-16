package middleware

import (
	"net/http"
	"strings"

	"forum-service/utils" // Or replace with "user-service/utils" if shared

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT access tokens and sets userID in context.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		userID, err := utils.ValidateAccessToken(tokenParts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Set userID in context for downstream handlers
		c.Set("userID", userID)
		c.Next()
	}
}
