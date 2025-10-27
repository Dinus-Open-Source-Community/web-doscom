package routes

import (
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBlogRoutes(rg *gin.RouterGroup, blogHandler *handler.BlogHandler) {
	rg.POST("/blogs", blogHandler.Create)
	rg.GET("/blogs", blogHandler.List)
	rg.GET("/blogs/:id", blogHandler.Get)
	rg.PUT("/blogs/:id", blogHandler.Update)
	rg.DELETE("/blogs/:id", blogHandler.Delete)
}
