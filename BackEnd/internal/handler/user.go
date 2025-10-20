package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Contoh handler untuk user

// CreateUserHandler untuk POST /api/user
func CreateUserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user": gin.H{
			"id":   1,
			"name": "Test User",
		},
	})
}

// GetUserHandler untuk GET /api/user/:id
func GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user",
		"user": gin.H{
			"id":   id,
			"name": "Test User",
		},
	})
}
