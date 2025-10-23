package auth

import (
	"log"
	"net/http"

	"web_doscom/internal/config"
	"web_doscom/internal/database/model"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return "", err
	}
	return hash, nil
}

func verifyPassword(password, hash string) bool {

	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Printf("Error verifying password: %v", err)
		return false
	}
	return match
}

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(app *config.Application) gin.HandlerFunc {
	return func(c *gin.Context) {

		// get the email and password from the req body
		var input loginRequest
		if c.Bind(&input) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read req body",
			})

			return
		}

		// look at the requested user
		var user model.User
		result := app.DB.First(&user, "email = ?", input.Email)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid email or Password",
			})
			return
		}
		// verify the password
		if !verifyPassword(input.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid email or password",
			})
			return
		}

		// generate token jwt
		token, err := Create_token(user.ID, user.Email, user.Username)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error: ": "Failed to create token",
			})

			return
		}

		// send back
		c.JSON(http.StatusOK, gin.H{
			"message:": "login success bolo, nasi padang satu bungkus",
			"token:":   token,
		})
	}
}
