package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// panic("failde to connect db : ")
		log.Fatal("failed to connect db : ", err)
	}

	DB = db
	fmt.Println("Success connect to db")

}
