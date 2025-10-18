package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type service struct {
	DB *gorm.DB
}

var dbInstance *service



func ConnectDB() *service {
	dbURL := os.Getenv("DBURL")
	fmt.Println("DBURL:", dbURL)
	// reuse connection
	if dbInstance != nil {
		return dbInstance
	}

	dsn := dbURL + "&search_path=" + os.Getenv("DB_TIMEZONE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect db : ", err)
	}

	dbInstance = &service{
		DB: db,
	}
	return dbInstance
}

// func (s *service) Close() error {
//  sqlDB, err := s.DB.DB()
//  if err != nil {
//      return err
//  }
//  log.Printf("Disconnected from database: %s", database)
//  return sqlDB.Close()
// }
