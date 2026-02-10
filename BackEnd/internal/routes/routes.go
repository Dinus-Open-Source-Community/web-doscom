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

		// Initialize storage service if MinIO is available
		var storageService *service.StorageService
		if app.MinioClient != nil {
			storageService = service.NewStorageService(app.MinioClient, app.DB)
		}

		galleryService := service.NewGalleryService(&app.Model.Gallery)
		galleryHandler := handler.NewUploadHandler(galleryService, storageService)

		pengurusService := service.NewPengurusService(&app.Model.Pengurus, galleryService)
		pengurusHandler := handler.NewPengurusHandler(pengurusService, storageService)

		blogService := service.NewBlogService(&app.Model.Blogs)
		blogHandler := handler.NewBlogHandler(blogService)

		// Register upload routes if MinIO is available
		if storageService != nil {
			storageHandler := handler.NewStorageHandler(storageService)
			routes.RegisterUploadRoutes(v1, storageHandler)
		}

		routes.AuthRoutes(v1, authHandler)
		routes.UserControllerKoor(v1, userHandler)
		routes.UserControllerRoute(v1, userHandler)
		routes.RegisterWorkRoutes(v1, workHandler)
		routes.GalleryRoute(v1, galleryHandler)
		routes.RegisterPengurusRoutes(v1, pengurusHandler)
		routes.RegisterBlogRoutes(v1, blogHandler)

		// swagger
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return g
}
