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

func Routes(g *gin.Engine, app *config.Application) http.Handler {

	v1 := g.Group("api/v1")
	{
		// auth
		authService := auth.NewAuthService(&app.Model.Users, &app.Model.RefreshToken)
		authHandler := auth.NewUserauth(authService)

		// user
		userService := service.NewUserService(&app.Model.Users)
		userHandler := handler.NewUserHandler(userService)

		// Initialize storage service if MinIO is available
		var storageService *service.StorageService
		if app.MinioClient != nil {
			// storageService = service.NewStorageService(app.MinioClient, app.DB)
			storageService = service.NewStorageService(
				app.MinioClient,
				&app.Model.FileUploads,
			)
		}
		// storageHandler := handler.NewStorageHandler(storageService)
		// Register upload routes if MinIO is available
		if storageService != nil {
			storageHandler := handler.NewStorageHandler(storageService)
			routes.RegisterUploadRoutes(v1, storageHandler)
		}

		galleryService := service.NewGalleryService(&app.Model.Gallery, storageService)
		galleryHandler := handler.NewUploadHandler(galleryService, storageService)

		// work
		workService := service.NewWorkService(
			&app.Model.Works,
			galleryService,
			&app.Model.Pengurus,
			&app.Model.WorkGallery,
		)
		workHandler := handler.NewWorkHandler(workService)
		// pengurus
		pengurusService := service.NewPengurusService(
			&app.Model.Pengurus,
			&app.Model.PengurusSosmed,
			galleryService,
			app.DB,
		)
		pengurusHandler := handler.NewPengurusHandler(pengurusService, storageService)

		blogService := service.NewBlogService(
			app.DB,
			&app.Model.Blogs,
			&app.Model.BlogGallery,
			galleryService,
		)
		blogHandler := handler.NewBlogHandler(blogService)

		routes.AuthRoutes(v1, authHandler)
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
