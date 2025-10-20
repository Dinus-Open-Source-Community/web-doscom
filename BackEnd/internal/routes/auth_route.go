package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/config"

	"github.com/gin-gonic/gin"
)

// app *server.Application
func AuthRoutes(r *gin.RouterGroup, app *config.Application) {

	r.POST("/login", auth.LoginHandler(app))

}
