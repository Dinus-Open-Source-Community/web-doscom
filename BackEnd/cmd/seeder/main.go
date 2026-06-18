package main

import (
	"log"

	"web_doscom/internal/config"
	env "web_doscom/internal/config"
	"web_doscom/internal/database"
	"web_doscom/internal/seeder"
)

func main() {
	env.LoadEnv()

	minioClient, err := config.InitMinioClient()
	if err != nil {
		log.Fatalf("Failed to initialize minio client: %v", err)
	}
	dbService := database.ConnectDB()
	if dbService == nil || dbService.DB == nil {
		log.Fatal("Failed to connect to the database")
	}

	log.Println("Database connection established for seeder.")

	seeder.RunAllSeeders(dbService.DB, minioClient)

	log.Println("Seeding process finished.")
}
