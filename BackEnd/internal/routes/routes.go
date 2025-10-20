package routes

import (
	"net/http"
	"web_doscom/internal/config"

	"github.com/gin-gonic/gin"
)

func Routes(app *config.Application) http.Handler {
	g := gin.Default()

	v1 := g.Group("api/v1")
	{
		AuthRoutes(v1, app)
	}
	return g
}
