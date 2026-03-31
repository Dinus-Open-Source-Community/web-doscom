package main

import (
	"log"
	"web_doscom/internal/auth"
	env "web_doscom/internal/config"
	"web_doscom/internal/database"
	"web_doscom/internal/database/model"
)

func main() {
	env.LoadEnv()

	// connect database
	db := database.ConnectDB()
	var user model.User

	hashPassword := auth.HashPassword("admin123")
	user = model.User{
		Username:  "super_admin",
		Email:     "superadmin@gmail.com",
		Password:  hashPassword,
		Role:      "Super_Admin",
		Full_name: "admin",
	}

	if err := db.DB.Create(&user).Error; err != nil {
		log.Fatal(err)
	}
}
