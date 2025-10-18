package env

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Try to load .env from BackEnd directory explicitly
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file from path: %v", err)
	}
}
