package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
	env "web_doscom/internal/config"
	"web_doscom/internal/database/model"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	Model        *model.UserModel
	RefreshToken *model.RefreshTokenModel
}

func NewAuthService(u *model.UserModel, m *model.RefreshTokenModel) *AuthService {
	return &AuthService{Model: u, RefreshToken: m}
}

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

var REFRESH_TOKEN_SIZE = 32

func Create_token(UserId int, email, username, role string) (string, error) {
	env.LoadEnv()

	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

	expirationTime := time.Now().Add(15 * time.Minute)

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

func GetCookie(r *http.Request, tokenName string) (string, error) {
	cookie, err := r.Cookie(tokenName)
	if err != nil {
		// switch {
		// case errors.Is(err, http.ErrNoCookie):
		// 	http.Error(w, "cookie not found", http.StatusUnauthorized)
		// default:
		// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// }
		return "", err
	}

	return cookie.Value, nil
}

func generateSecureToken() (string, error) {
	bytes := make([]byte, REFRESH_TOKEN_SIZE)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(bytes), nil
}

func generateRefreshToken(userId int) (*model.RefreshToken, string, error) {
	tokenString, err := generateSecureToken()
	if err != nil {
		return nil, "", err
	}
	tokenHash := HashPassword(tokenString)
	expiredAt := time.Now().Add(5 * 24 * time.Hour)

	refreshToken := &model.RefreshToken{
		UserId:    userId,
		Token:     tokenHash,
		Expires:   &expiredAt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return refreshToken, tokenHash, nil
}

func (h *AuthService) validateRefreshToken(tokenString string) (string, error) {
	// take refresh token from database
	refreshTokenHash := HashPassword(tokenString)
	refreshToken, err := h.RefreshToken.GetRefreshToken(refreshTokenHash)
	if err != nil {
		return "", err
	}

	// check expires refresh token
	if time.Now().After(*refreshToken.Expires) {
		return "", errors.New("Refresh token expires")
	}

	// check user_id
	user, err := h.Model.GetUserById(refreshToken.UserId)
	if err != nil {
		return "", err
	}

	// create new access_token
	accessToken, err := Create_token(user.ID, user.Email, user.Username, user.Role)
	if err != nil {
		return "", err
	}

	// refreshToken.UpdatedAt = time.Now()
	// if err := h.RefreshToken.UpdatePartial(refreshToken.Token, map[string]any{
	// 	"token_hash": refreshToken.UpdatedAt,
	// }); err != nil {
	// 	return "", err
	// }

	return accessToken, nil
}

// wrapper model user
func (h *AuthService) FindByEmail(email string) (*model.User, error) {
	return h.Model.FindByEmail(email)
}

func (h *AuthService) InsertUser(user *model.User) error {
	return h.Model.InsertUser(user)
}

// wrapper model refreshToken
func (h *AuthService) CreateRefreshToken(refreshToken *model.RefreshToken) error {
	return h.RefreshToken.CreateRefreshToken(refreshToken)
}
