package routes

import (
	"user-service/controllers"
	"user-service/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterAuthRoutes defines all auth-related routes and protected API routes
func RegisterAuthRoutes(r *gin.Engine) {
	// Public auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register) // Handles user registration
		auth.POST("/login", controllers.Login)       // Handles user login
		auth.POST("/refresh", controllers.Refresh)   // Handles token refresh
		auth.POST("/logout", controllers.Logout)     // Handles logout
	}

	// Protected API routes (requires valid token)
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware()) // Middleware validates JWT
	{
		api.GET("/profile", func(c *gin.Context) {
			userID := c.MustGet("userID").(uint) // Extracted from token
			c.JSON(200, gin.H{
				"message": "Authenticated",
				"user_id": userID,
			})
		})
	}
}
