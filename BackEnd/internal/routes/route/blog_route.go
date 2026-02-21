package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(rg *gin.RouterGroup, blogHandler *handler.BlogHandler) {
	blogRoutes := rg.Group("/blogs")
	blogRoutes.Use(auth.AuthMiddleware("Super_Admin", "Kor_Medcrev"))
	{
		blogRoutes.POST("/", blogHandler.CreateBlog) // create new blog
		blogRoutes.GET("/:id", blogHandler.Get)      // get blog by id
		blogRoutes.GET("", blogHandler.List)         // get all blog
		blogRoutes.PUT("/:id", blogHandler.Update)   // update blog by id
		blogRoutes.PATCH("/:id", blogHandler.Update) // update blog by id
		blogRoutes.PUT("/:id/kategori", blogHandler.UpdateKategori)
		blogRoutes.PATCH("/:id/kategori", blogHandler.UpdateKategori)
		blogRoutes.GET("/kategori/:kategori", blogHandler.ListByKategori)
		blogRoutes.DELETE("/:id", blogHandler.Delete)
	}
}
