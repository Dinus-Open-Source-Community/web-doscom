package server

import (
	"net/http"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up example routes for the Gin engine
func RegisterRoutes(r *gin.Engine) {

	// TEST PING ROUTE
	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// ROUTE TO GET USER BY ID
	r.GET("/api/user/:id", handler.GetUserHandler)

	// ROUTE TO CREATE USER
	r.POST("/api/user", handler.CreateUserHandler)
}

// Routes returns the Gin engine with all routes registered
func (app *Application) Routes() *gin.Engine {
	g := gin.Default()
	RegisterRoutes(g)
	return g
}
