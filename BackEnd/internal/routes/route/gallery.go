package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func GalleryRoute(r *gin.RouterGroup, GalleryHandler *handler.GalleryHandler) {
	crevmed := r.Group("/gallery")
	crevmed.Use(auth.AuthMiddleware("Kor_Medcrev"))
	{
		crevmed.POST("/", GalleryHandler.InsertGallery) // insert gallery
		// get gallery by id
		// get all gallery
		// update gallery by id
		// delete gallery by id
	}
}
