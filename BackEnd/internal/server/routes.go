package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up example routes for the Gin engine
func RegisterRoutes(r *gin.Engine) {

	// TEST PING ROUTE
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// ROUTE TO CREATE USER

	r.GET("/api/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "Get user",
			"user": gin.H{
				"id":   id,
				"name": "Test User",
			},
		})
	})

	// ROUTE TO CREATE USER
	r.POST("/api/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
			"user": gin.H{
				"id":   1,
				"name": "Test User",
			},
		})
	})
}

// Routes returns the Gin engine with all routes registered
func (app *Application) Routes() *gin.Engine {
	g := gin.Default()
	RegisterRoutes(g)
	return g
}
