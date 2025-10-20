package auth

import (
	"os"
	"time"
	env "web_doscom/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId   int    `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"full_name"`
	jwt.RegisteredClaims
}

func Create_token(UserId int, email string, username string) (string, error) {
	env.LoadEnv()

	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		UserId:   UserId,
		Email:    email,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "doscom-backend", // optional, untuk identifikasi pembuat token
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenstring, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenstring, nil
}
