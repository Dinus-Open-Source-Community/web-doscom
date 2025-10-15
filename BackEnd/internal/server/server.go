package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"web_doscom/internal/database"

	"gorm.io/gorm"
)

type Application struct {
	Port  int
	DB    *gorm.DB
	Model database.Models
}

func (app *Application) NewServer() error {

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("🚀 Server running on Port %d", app.Port)
	return server.ListenAndServe()

}
