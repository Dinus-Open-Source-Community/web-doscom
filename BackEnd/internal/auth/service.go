package auth

import (
	"log"
	"os"
	"time"
	env "web_doscom/internal/config"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId   int    `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"full_name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func Create_token(UserId int, email, username, role string) (string, error) {
	env.LoadEnv()

	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		UserId:   UserId,
		Email:    email,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "doscom-backend", // optional, untuk identifikasi pembuat token
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenstring, nil
}

func HashPassword(password string) string {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return ""
	}
	return hash
}

func verifyPassword(password, hash string) bool {

	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Printf("Error verifying password: %v", err)
		return false
	}
	return match
}
