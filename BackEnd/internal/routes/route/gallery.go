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
		crevmed.GET("/", GalleryHandler.GetGalleryByType) // get gallery by id
	}

	// private api -> need auth
	crevmedAuth := crevmed.Group("/")
	crevmedAuth.Use(auth.AuthMiddleware("Kor_Medcrev"))
	{
		crevmed.POST("/", GalleryHandler.InsertGallery) // insert gallery
		// get all gallery
		crevmed.DELETE("/:id", GalleryHandler.DeleteGallery) // update gallery by id
		// delete gallery by id
	}
}
