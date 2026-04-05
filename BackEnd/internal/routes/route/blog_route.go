package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(rg *gin.RouterGroup, blogHandler *handler.BlogHandler) {

	publickRoutes := rg.Group("/blogs")
	{
		publickRoutes.GET("", blogHandler.ListByKategori)
		publickRoutes.GET("/:id", blogHandler.GetBlogByID)
	}

	privateRoutes := rg.Group("/admin/blogs")
	privateRoutes.Use(auth.AuthMiddleware("SuperAdmin", "KoorMedcrev"))
	{
		privateRoutes.GET("", blogHandler.ListBlogs)
		privateRoutes.POST("", blogHandler.CreateBlog)
		privateRoutes.GET("/:id", blogHandler.GetBlogByID)
		privateRoutes.PUT("/:id", blogHandler.Update)
		privateRoutes.DELETE("/:id", blogHandler.Delete)
	}
}
