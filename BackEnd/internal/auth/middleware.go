package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateAuth(tokenString string) (*Claims, error) {
	claims := &Claims{}
	// parse the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method mbot: %v", token.Header["alg"])
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

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error:": "Authorization is missing",
			})
			c.Abort()
			return
		}

		// format
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error:": "invalid token format",
			})
			c.Abort()
			return
		}

		tokenstring := tokenParts[1]
		// validasi token
		claims, err := ValidateAuth(tokenstring)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error:": err.Error(),
			})

			c.Abort()
			return
		}

		c.Set("user_id", claims.UserId)
		c.Set("email", claims.Email)
		c.Set("username", claims.Username)

		c.Next()
	}
}
