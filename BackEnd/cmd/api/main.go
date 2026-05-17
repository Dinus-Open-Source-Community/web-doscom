package main

import (
	"fmt"

	"log"
	"os"
	"strconv"
	"web_doscom/internal/config"
	env "web_doscom/internal/config"
	"web_doscom/internal/database"
	"web_doscom/internal/server"
	"web_doscom/internal/utils"

	_ "web_doscom/docs"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Test user creation API handler
// @title Web Doscom API
// @version 1.0
// @description API Documentation Web Doscom
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// load env
	env.LoadEnv()

	// load costum binding
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("socialurl", utils.ValidateURL)
	}

	// connect database
	db := database.ConnectDB()
	fmt.Println("DBURL:", os.Getenv("DBURL"))

	models := database.NewModel(db.DB)
	portstr := os.Getenv("PORT")
	port, err := strconv.Atoi(portstr)
	if err != nil {
		log.Fatal("Invalid port value")
	}

	// Initialize MinIO client
	minioClient, err := config.InitMinioClient()
	if err != nil {
		log.Printf("Warning: MinIO initialization failed: %v", err)
		log.Println("File upload features will be disabled")
	} else {
		log.Println("MinIO client initialized successfully")
	}

	app := &config.Application{
		Port:        port,
		DB:          db.DB,
		Model:       models,
		MinioClient: minioClient,
	}

	if err := server.NewServer(app); err != nil {
		log.Fatal(err)
	}

}
