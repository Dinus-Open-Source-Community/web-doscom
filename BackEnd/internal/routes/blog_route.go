package routes

import (
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(rg *gin.RouterGroup, blogHandler *handler.BlogHandler) {
	blogRoutes := rg.Group("/blogs")
	{
		blogRoutes.POST("", blogHandler.Create)
		blogRoutes.GET("/:id", blogHandler.Get)
		blogRoutes.GET("", blogHandler.List)
		blogRoutes.PUT("/:id", blogHandler.Update)
		blogRoutes.PATCH("/:id", blogHandler.Update)
		blogRoutes.PUT("/:kategori/:id", blogHandler.UpdateKategori)
		blogRoutes.PATCH("/:kategori/:id", blogHandler.UpdateKategori)
		blogRoutes.GET("/kategori/:kategori", blogHandler.ListByKategori)
		blogRoutes.DELETE("/:id", blogHandler.Delete)
	}
}
