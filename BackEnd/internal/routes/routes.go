package routes

import (
	"net/http"
	"web_doscom/internal/config"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func Routes(app *config.Application) http.Handler {
	g := gin.Default()

	// v1 := g.Group("api/v1")
	v1 := g.Group("/api/v1")
	AuthRoutes(v1, app)

	workHandler := handler.NewWorkHandler(app.DB)
	RegisterWorkRoutesV2(v1, workHandler)

	activitiesHandler := handler.NewActivitiesHandler(app.DB)
	RegisterActivitiesRoutesV1(v1, activitiesHandler)

	// RegisterBlogRoutes(g, app.DB)
	blogHandler := handler.NewBlogHandler(app.DB)
	RegisterBlogRoutes(v1, blogHandler)

	// Tambahkan route lain di sini
	return g
}
