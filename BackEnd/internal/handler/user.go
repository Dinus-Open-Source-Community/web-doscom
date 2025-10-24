package handler

import (
	"net/http"
	"web_doscom/internal/database/model"

	"github.com/gin-gonic/gin"
)

// CreateUserHandler untuk POST /api/user
func CreateUserHandler(UserModel *model.UserModel) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input model.RegisterRequest

		if c.Bind(&input) != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read req body",
			})

			return
		}

		user := &model.User{
			Username:  input.Username,
			Email:     input.Email,
			Role:      input.Role,
			Password:  input.Password,
			Full_name: input.Fullname,
		}

		if err := UserModel.InsertUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create user",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
		})
	}
}

// GetUserHandler untuk GET /api/user/:id
func GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user",
		"user": gin.H{
			"id":   id,
			"name": "Test User",
		},
	})
}
