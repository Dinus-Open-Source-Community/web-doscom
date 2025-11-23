package routes

import (
	"net/http"
	"web_doscom/internal/auth"
	"web_doscom/internal/config"
	"web_doscom/internal/handler"
	routes "web_doscom/internal/routes/route"
	"web_doscom/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(app *config.Application) http.Handler {
	g := gin.Default()

	v1 := g.Group("api/v1")
	// Tambahkan route lain di sini
	{

		authHandler := auth.NewUserauth(&app.Model.Users)
		workHandler := handler.NewWorkHandler(&app.Model.Works)
		userHandler := handler.NewUserHandler(&app.Model.Users)

		galleryService := service.NewGalleryService(&app.Model.Gallery)
		galleryHandler := handler.NewUploadHandler(galleryService)

		pengurusService := service.NewPengurusService(&app.Model.Pengurus, galleryService)
		pengurusHandler := handler.NewPengurusHandler(pengurusService)

		routes.AuthRoutes(v1, authHandler)
		routes.UserControllerKoor(v1, userHandler)
		routes.UserControllerRoute(v1, userHandler)
		routes.RegisterWorkRoutes(v1, workHandler)
		routes.GalleryRoute(v1, galleryHandler)
		routes.RegisterPengurusRoutes(v1, pengurusHandler)

		// swagger
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return g
}
