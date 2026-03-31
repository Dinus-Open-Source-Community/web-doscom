package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func GalleryRoute(r *gin.RouterGroup, GalleryHandler *handler.GalleryHandler) {
	publicRoutes := r.Group("/gallery")
	// public api
	{
<<<<<<< HEAD
		// get all gallery and gallery by filter -> year
		publicRoutes.GET("/", GalleryHandler.GetAllGalleryByYear)

	}

	// private api -> need auth
	privateRoutes := r.Group("/admin/gallery")
	privateRoutes.Use(auth.AuthMiddleware("SuperAdmin", "KoorMedcrev"))
	{
		// insert gallery
		privateRoutes.POST("/", GalleryHandler.InsertGallery)
		// delete gallery by id
		privateRoutes.DELETE("/:id", GalleryHandler.DeleteGallery)
=======
		crevmed.GET("/", GalleryHandler.GetGalleryByType) // get gallery by type
	}

	// private api -> need auth
	crevmedAuth := crevmed.Group("/")
	crevmedAuth.Use(auth.AuthMiddleware("Kor_Medcrev", "Super_Admin"))
	{
		crevmedAuth.POST("/", GalleryHandler.InsertGallery)      // insert gallery
		crevmedAuth.DELETE("/:id", GalleryHandler.DeleteGallery) // delete gallery by id
>>>>>>> master
	}
}
