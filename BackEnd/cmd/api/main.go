package main

import (
	"log"
	"os"
	"strconv"
	"web_doscom/internal/config"
	env "web_doscom/internal/config"
	"web_doscom/internal/database"
	"web_doscom/internal/server"
)

func main() {
	// load env
	env.LoadEnv()

	// connect database
	db := database.ConnectDB()

	models := database.NewModel(db.DB)
	portstr := os.Getenv("PORT")
	port, err := strconv.Atoi(portstr)
	if err != nil {
		log.Fatal("Invalid port value")
	}
	app := &config.Application{
		Port:  port,
		DB:    db.DB,
		Model: models,
	}

	if err := server.NewServer(app); err != nil {
		log.Fatal(err)
	}

}
