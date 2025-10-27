package routes

import (
	"web_doscom/internal/auth"

	"github.com/gin-gonic/gin"
)

// app *server.Application
func AuthRoutes(r *gin.RouterGroup, AuthHandler *auth.AuthHandler) {

	r.POST("/login", AuthHandler.LoginHandler)    // login route for admin and superadmin
	r.POST("/register", AuthHandler.RegisterUser) // register route for admin and superadmin

}
