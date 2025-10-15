package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Service interface {
// 	Close() error
// }

type service struct {
	DB *gorm.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	username   = os.Getenv("DB_USER")
	password   = os.Getenv("DB_PASSWORD")
	host       = os.Getenv("DB_HOST")
	port       = os.Getenv("DB_PORT")
	timezone   = os.Getenv("DB_TIMEZONE")
	dbInstance *service
)

func ConnectDB() *service {
	// reuse connection
	if dbInstance != nil {
		return dbInstance
	}

	var err error
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, timezone)
	// dsn := os.Getenv("DB")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// panic("failde to connect db : ")
		log.Fatal("failed to connect db : ", err)
	}

	dbInstance = &service{
		DB: db,
	}

	return dbInstance

}

// func (s *service) Close() error {
// 	sqlDB, err := s.DB.DB()
// 	if err != nil {
// 		return err
// 	}
// 	log.Printf("Disconnected from database: %s", database)
// 	return sqlDB.Close()
// }
