package auth

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"
	env "web_doscom/internal/config"

	"github.com/alexedwards/argon2id"
	//"github.com/docker/docker/daemon/names"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	//"golang.org/x/tools/go/analysis/passes/stringintconv"
)

type Claims struct {
	UserId   int    `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"full_name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type Cookies struct {
	Name     string
	Value    string
	Path     string
	Expires  time.Time
	Secure   bool
	HttpOnly bool
	SameSite http.SameSite
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

func SetCustomCookie(c *gin.Context, cfg Cookies) {
	cookie := &http.Cookie{
		Name:     cfg.Name,
		Value:    cfg.Value,
		Path:     cfg.Path,
		Expires:  cfg.Expires,
		Secure:   cfg.Secure,
		HttpOnly: cfg.HttpOnly,
		SameSite: cfg.SameSite,
	}

	http.SetCookie(c.Writer, cookie)
}

func GetCookie(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("AccessToken")
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusUnauthorized)
		default:
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return "", err
	}

	return cookie.Value, nil
}
