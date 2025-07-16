package main

import (
	"log"

	"user-service/database" // ✅ Add this
	"user-service/routes"
	"user-service/utils"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	// ✅ Initialize database
	database.InitDB()

	// ✅ Initialize Redis
	utils.InitRedis()

	// ✅ Create Gin router
	router := gin.Default()

	// ✅ Add this line to fix the CORS error
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5500", "http://127.0.0.1:5500"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// ✅ Register auth routes
	routes.RegisterAuthRoutes(router)

	// ✅ Test route
	router.GET("/test-redis", routes.RedisHandler)

	// ✅ Run the server
	log.Println("Server running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
