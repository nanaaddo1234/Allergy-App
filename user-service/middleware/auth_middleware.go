package middleware

import (
	"net/http"
	"strings"

	"user-service/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a Gin middleware that validates JWT access tokens.
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

		// Store the user ID in the context for use in handlers
		c.Set("userID", userID)
		c.Next()
	}
}
