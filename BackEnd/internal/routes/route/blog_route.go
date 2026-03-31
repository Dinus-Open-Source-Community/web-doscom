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
<<<<<<< HEAD
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
=======
		blogRoutes.POST("", blogHandler.Create)      // create new blog
		blogRoutes.GET("/:id", blogHandler.Get)      // get blog by id
		blogRoutes.GET("", blogHandler.List)         // get all blog
		blogRoutes.PUT("/:id", blogHandler.Update)   // update blog by id
		blogRoutes.PATCH("/:id", blogHandler.Update) // update blog by id
		blogRoutes.GET("/kategori/:kategori", blogHandler.ListByKategori)
		blogRoutes.DELETE("/:id", blogHandler.Delete)
>>>>>>> master
	}
}
