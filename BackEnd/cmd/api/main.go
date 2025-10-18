package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"web_doscom/internal/database"
	"web_doscom/internal/env"
	"web_doscom/internal/server"
)

// Test user creation API handler
func main() {
	// load env
	env.LoadEnv()

	// connect database
	db := database.ConnectDB()
	fmt.Println("DBURL:", os.Getenv("DBURL"))

	models := database.NewModel(db.DB)
	portstr := os.Getenv("PORT")
	port, err := strconv.Atoi(portstr)
	if err != nil {
		log.Fatal("Invalid port value")
	}
	app := &server.Application{
		Port:  port,
		DB:    db.DB,
		Model: models,
	}

	// Use Gin router from server.Routes()
	fmt.Printf("Gin API server running on :%d\n", port)
	log.Fatal(app.Routes().Run(fmt.Sprintf(":%d", port)))
}
