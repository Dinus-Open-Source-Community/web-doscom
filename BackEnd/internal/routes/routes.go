package routes

import (
	"net/http"
	"web_doscom/internal/config"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func Routes(app *config.Application) http.Handler {
	g := gin.Default()

	v1 := g.Group("api/v1")
	AuthRoutes(v1, app)

	workHandler := handler.NewWorkHandler(app.DB)
	RegisterWorkRoutesV2(v1, workHandler)

	// Tambahkan route lain di sini
	return g
}
