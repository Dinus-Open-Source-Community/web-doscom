package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"web_doscom/internal/config"
	"web_doscom/internal/routes"

	"github.com/gin-gonic/gin"
)

func NewServer(app *config.Application) error {

	g := gin.Default()

	// CORS middleware
	g.Use(CORSMiddleware())

	// routes
	routes.Routes(g, app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Port),
		Handler:      g,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("🚀 Server running on Port %d", app.Port)
	return server.ListenAndServe()

}
