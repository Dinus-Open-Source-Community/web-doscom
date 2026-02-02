package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func GalleryRoute(r *gin.RouterGroup, GalleryHandler *handler.GalleryHandler) {
	crevmed := r.Group("/gallery")
	// public api
	{
		crevmed.GET("/", GalleryHandler.GetGalleryByType) // get gallery by type
	}

	// private api -> need auth
	crevmedAuth := crevmed.Group("/")
	crevmedAuth.Use(auth.AuthMiddleware("Kor_Medcrev", "Super_Admin"))
	{
		crevmedAuth.POST("/", GalleryHandler.InsertGallery)      // insert gallery
		crevmedAuth.DELETE("/:id", GalleryHandler.DeleteGallery) // delete gallery by id
	}
}
