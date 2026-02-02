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
		crevmed.POST("/", GalleryHandler.InsertGallery)      // insert gallery
		crevmed.GET("/all", GalleryHandler.GetAllGallery)    // get all gallery
		crevmed.GET("/:id", GalleryHandler.GetGalleryByID)   // get gallery by id
		crevmed.PUT("/:id", GalleryHandler.UpdateGallery)    // update gallery by id
		crevmed.DELETE("/:id", GalleryHandler.DeleteGallery) // delete gallery by id
	}
}
