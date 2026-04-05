package routes

import (
	"web_doscom/internal/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, AuthHandler *auth.AuthHandler) {

	r.POST("/login", AuthHandler.LoginHandler)
	r.POST("/register", AuthHandler.RegisterUser)
	r.POST("/auth/refresh", AuthHandler.RefreshToken)
	r.POST("/auth/logout", AuthHandler.Logout)

}
