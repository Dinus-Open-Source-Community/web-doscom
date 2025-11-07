package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"slices"

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

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// get cookie and token
		tokenString, err := c.Cookie("AccessToken")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Cookie not found",
			})
			c.Abort()
			return
		}

		// validasi token
		claims, err := ValidateAuth(tokenString)
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
		c.Set("role", claims.Role)

		// cek role
		// allowedRoles := []string{"Super_Admin", "Kor_Pemro", "Kor_Jaringan", "Kor_Data", "Kor_Medcrev", "BPH", "pemro_ang", "jaringan_ang", "medcrev_ang", "data_ang", "BPH_ang"}

		if len(allowedRoles) > 0 {
			if !slices.Contains(allowedRoles, claims.Role) {
				c.JSON(http.StatusForbidden, gin.H{
					"error": "forbidden",
				})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
