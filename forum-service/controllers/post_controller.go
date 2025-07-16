package controllers

import (
	"forum-service/database"
	"forum-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePost handles POST /api/posts
func CreatePost(c *gin.Context) {
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		UserID  uint   `json:"user_id" binding:"required"`
		Tags    string `json:"tags"` // comma-separated
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  input.UserID,
		// models.Post.Tags should be of type string to store comma-separated tags
		Votes: 0,
	}

	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save post"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetAllPosts handles GET /api/posts
func GetAllPosts(c *gin.Context) {
	var posts []models.Post
	if err := database.DB.Order("votes desc").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// PostResponse is used to send tags as a slice
type PostResponse struct {
	ID      uint   `json:"ID"`
	Title   string `json:"Title"`
	Content string `json:"Content"`

	UserID uint `json:"UserID"`
}

// GetPostByID handles GET /api/posts/:id
func GetPostByID(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	response := PostResponse{
		ID:      post.ID,
		Title:   post.Title,
		Content: post.Content,

		UserID: post.UserID,
	}

	c.JSON(http.StatusOK, response)
}
