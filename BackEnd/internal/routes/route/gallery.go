package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func GalleryRoute(r *gin.RouterGroup, GalleryHandler *handler.GalleryHandler) {
	publicRoutes := r.Group("/gallery")
	{
		publicRoutes.GET("", GalleryHandler.GetAllGalleryAndByYear)

	}

	privateRoutes := r.Group("/admin/gallery")
	privateRoutes.Use(auth.AuthMiddleware("SuperAdmin", "KoorMedcrev"))
	{
		privateRoutes.POST("", GalleryHandler.InsertGallery)
		privateRoutes.DELETE("/:id", GalleryHandler.DeleteGallery)
	}
}
