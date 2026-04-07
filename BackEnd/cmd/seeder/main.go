package main

import (
	"log"
	
	env "web_doscom/internal/config"
	"web_doscom/internal/database"
	"web_doscom/internal/seeder"
)

func main() {
	// Load environment variables
	env.LoadEnv()

	// Initialize database connection
	dbService := database.ConnectDB()
	if dbService == nil || dbService.DB == nil {
		log.Fatal("Failed to connect to the database")
	}

	log.Println("Database connection established for seeder.")

	// Execute all seeders
	seeder.RunAllSeeders(dbService.DB)

	log.Println("Seeding process finished.")
}
