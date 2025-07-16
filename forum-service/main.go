package main

import (
	"fmt"
	"os"

	"forum-service/database"
	"forum-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// ✅ Initialize database
	database.InitDB()

	// ✅ Create Gin router
	router := gin.Default()

	// ✅ Configure CORS middleware BEFORE registering routes
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5500", "http://127.0.0.1:5500"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// ✅ Register forum routes AFTER CORS middleware is applied
	routes.RegisterForumRoutes(router)

	// ✅ Set port (default to 8081 if not set)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	fmt.Printf("Forum service is running on http://localhost:%s\n", port)

	// ✅ Start the server
	if err := router.Run(":" + port); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
