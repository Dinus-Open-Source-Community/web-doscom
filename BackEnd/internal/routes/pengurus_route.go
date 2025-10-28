package routes

import (
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPengurusRoutesV2(rg *gin.RouterGroup, pengurusHandler *handler.PengurusHandler) {

	pengurus := rg.Group("/pengurus")
	{
		pengurus.POST("", pengurusHandler.CreatePengurus)
		pengurus.GET("/:id", pengurusHandler.GetPengurus)
		pengurus.PUT("/:id", pengurusHandler.UpdatePengurus)
		pengurus.DELETE("/:id", pengurusHandler.DeletePengurus)
	}
}
