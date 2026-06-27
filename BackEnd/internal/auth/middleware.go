package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"slices"
	"web_doscom/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var RoleGroups = map[string][]string{
	"KOOR":     {"KoorPemro", "KoorJaringan", "KoorData", "KoorMedcrev", "BPH"},
	"ADMIN":    {"SuperAdmin"},
	"BPH":      {"BPH"},
	"PENGURUS": {"pemroAnggota", "jaringanAnggota", "medcrevAnggota", "dataAnggota", "BPHAnggota"},
}

func ValidateAuth(tokenString string) (*Claims, error) {
	claims := &Claims{}
	// parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("Token Expired")
		}
		return nil, fmt.Errorf("Invalid Token: %w", err)
	}

	// take claims
	if !token.Valid {
		return nil, fmt.Errorf("Invalid token bro")
	}

	return claims, nil
}

// Assuming Verify_token is a new or renamed function that returns *jwt.Token
// and handles the parsing and validation.
// For the purpose of this edit, I will define a placeholder for Verify_token
// based on the logic implied by the AuthMiddleware change.
func Verify_token(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("Token Expired")
		}
		return nil, fmt.Errorf("Invalid Token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return token, nil
}

func isRoleAllowed(User_role string, allowedRoles []string) bool {
	for _, role := range allowedRoles {
		if roles, exist := RoleGroups[role]; exist {
			if slices.Contains(roles, User_role) {
				return true
			}
		} else {
			if User_role == role {
				return true
			}
		}
	}
	return false
}

func getAccessToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
			return "", fmt.Errorf("Authroization header must be bearer token")
		}
		return authHeader[7:], nil
	}

	accessToken, err := c.Cookie("AccessToken")
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// get access token
		accessToken, err := getAccessToken(c)
		if err != nil {
			response.Error(
				c,
				http.StatusUnauthorized,
				"token not found",
				err,
			)
			c.Abort()
			return
		}
		// Validate token
		claims, err := ValidateAuth(accessToken)
		if err != nil {
			response.Error(
				c,
				http.StatusUnauthorized,
				"Invalid token",
				err,
			)
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", claims.UserId)
		c.Set("role", claims.Role)

		// Check role if specified
		if len(allowedRoles) > 0 {
			if !isRoleAllowed(claims.Role, allowedRoles) {
				response.Error(
					c,
					http.StatusForbidden,
					"you are not allowed access this resource",
					nil,
				)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
