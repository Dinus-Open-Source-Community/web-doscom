package auth

import (
	"log"
	"net/http"

	"web_doscom/internal/config"
	"web_doscom/internal/database/model"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
)

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
				"error": "Invalid email ",
			})
			return
		}
		// veerify the password
		if !verifyPassword(input.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":    "Invalid  password",
				"pw":       user.Password,
				"pw input": input.Password,
			})
			return
		}

		// check the role
		if user.Role != "Admin" && user.Role != "Super_Admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Acces denied, users have no acces",
			})

			return
		}

		// generate token jwt
		token, err := Create_token(user.Id, user.Email, user.Username, user.Role)

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

func RegisterUser(app *config.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the request body
		var input model.RegisterRequest
		if c.Bind(&input) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read req body",
			})

			return
		}

		// hash the password
		passwordHash := HashPassword(input.Password)
		log.Printf("Role value: '%s'", input.Role)

		// mapping data user
		user := model.User{
			Username:  input.Username,
			Email:     input.Email,
			Password:  passwordHash,
			Role:      input.Role,
			Full_name: input.Fullname,
		}

		// insert to database
		err := app.Model.Users.InsertUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to register user",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "user created Succesfully",
		})

	}
}
