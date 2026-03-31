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

var RoleGroups = map[string][]string{
	"KOOR":    {"KoorPemro", "KoorJaringan", "KoorData", "KoorMedcrev", "BPH"},
	"ADMIN":   {"SuperAdmin"},
	"BPH":     {"BPH"},
	"ANGGOTA": {"pemroAnggota", "jaringanAnggota", "medcrevAnggota", "dataAnggota", "BPHAnggota"},
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

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string

		// Try to get token from Authorization header first (Bearer token)
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			// Check if it's a Bearer token
			if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
				tokenString = authHeader[7:]
			} else {
				// If Authorization header exists but not Bearer format, use it as-is
				tokenString = authHeader
			}
		}

		// If no Authorization header, try to get token from cookie
		if tokenString == "" {
			var err error
<<<<<<< HEAD
			tokenString, err = c.Cookie("RefreshToken")
=======
			tokenString, err = c.Cookie("AccessToken")
>>>>>>> master
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Authentication required. Please provide a valid token in Authorization header or cookie",
				})
				c.Abort()
				return
			}
		}

		// Validate token
		claims, err := ValidateAuth(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		// Set user info in context
		c.Set("user_id", claims.UserId)
		// c.Set("email", claims.Email)
		// c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		// Check role if specified
		if len(allowedRoles) > 0 {
			if !isRoleAllowed(claims.Role, allowedRoles) {
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

// helper function for check role
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
