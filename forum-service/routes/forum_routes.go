package routes

import (
	"forum-service/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterForumRoutes sets up all /api/posts routes
func RegisterForumRoutes(r *gin.Engine) {
	posts := r.Group("/api/posts")
	{
		posts.POST("/", controllers.CreatePost)
		posts.GET("/", controllers.GetAllPosts)
		posts.GET("/:id", controllers.GetPostByID)
	}
}
