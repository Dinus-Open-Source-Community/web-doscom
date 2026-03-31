package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(rg *gin.RouterGroup, blogHandler *handler.BlogHandler) {

	// public routes
	publickRoutes := rg.Group("/blogs")
	{
		// get blog by id
		publickRoutes.GET("/:id", blogHandler.GetBlogByID)
		// get all blog and blog by kategori
		publickRoutes.GET("/", blogHandler.ListByKategori)
	}

	// private routes
	privateRoutes := rg.Group("/admin/blogs")
	privateRoutes.Use(auth.AuthMiddleware("SuperAdmin", "KoorMedcrev"))
	{
		// list blog by kategori
		privateRoutes.GET("", blogHandler.ListBlogs)
		// get blog by id
		privateRoutes.GET("/:id", blogHandler.GetBlogByID)
		// create new blog
		privateRoutes.POST("", blogHandler.CreateBlog)
		// update blog by id
		privateRoutes.PUT("/:id", blogHandler.Update)
		// delete blog by id
		privateRoutes.DELETE("/:id", blogHandler.Delete)
	}
}
