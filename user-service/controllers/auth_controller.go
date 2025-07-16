package controllers

import (
	"net/http"
	"strconv"
	"user-service/database"
	"user-service/models"
	"user-service/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := utils.HashPassword(input.Password)

	user := models.User{
		Email:        input.Email,
		Username:     input.Username,
		PasswordHash: hashedPassword,
		Role:         "patient",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	refreshToken, _ := utils.CreateRefreshToken(user.ID)
	c.SetCookie("refresh_token", refreshToken, 7*24*60*60, "/", "localhost", true, true)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Refresh(c *gin.Context) {
	oldToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing refresh token"})
		return
	}

	userIDStr, err := utils.GetUserIDFromToken(oldToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	userID, _ := strconv.Atoi(userIDStr)
	token, _ := utils.GenerateJWT(uint(userID), "patient")
	newToken, _ := utils.RotateRefreshToken(oldToken, uint(userID))

	c.SetCookie("refresh_token", newToken, 7*24*60*60, "/", "localhost", true, true)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Logout(c *gin.Context) {
	token, err := c.Cookie("refresh_token")
	if err == nil {
		utils.DeleteRefreshToken(token)
	}
	c.SetCookie("refresh_token", "", -1, "/", "localhost", true, true)
	c.Status(http.StatusNoContent)
}
