package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPengurusRoutes(rg *gin.RouterGroup, pengurusHandler *handler.PengurusHandler) {
	pengurus := rg.Group("/pengurus")
	pengurus.Use(auth.AuthMiddleware("ANGGOTA", "KOOR", "BPH", "ADMIN"))
	{
		pengurus.POST("/", pengurusHandler.CreatePengurus)
		pengurus.GET("/:id", pengurusHandler.GetPengurus)
		pengurus.PUT("/:id", pengurusHandler.UpdatePengurus)
		pengurus.GET("/", pengurusHandler.GetAllPengurus)
		pengurus.DELETE("/:id", pengurusHandler.DeletePengurus)
	}

}
