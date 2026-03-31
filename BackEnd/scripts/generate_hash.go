package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// func main() {
// 	password := "admin123"
//
// 	// Generate bcrypt hash
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Printf("Password: %s\n", password)
// 	fmt.Printf("Bcrypt Hash: %s\n", string(hash))
// }
