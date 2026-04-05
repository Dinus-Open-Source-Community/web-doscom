package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPengurusRoutes(rg *gin.RouterGroup, pengurusHandler *handler.PengurusHandler) {

	publicRoutes := rg.Group("/pengurus")
	{
		publicRoutes.GET("/division/:division", pengurusHandler.GetAllPengurus)
	}

	privateRoutes := rg.Group("/admin/pengurus")
	privateRoutes.Use(auth.AuthMiddleware("ANGGOTA", "KOOR", "BPH", "ADMIN"))
	{
		privateRoutes.POST("", pengurusHandler.CreatePengurus)
		privateRoutes.GET("/:id", pengurusHandler.GetPengurusByID)
		privateRoutes.PUT("/:id", pengurusHandler.UpdatePengurus)

		restricted := privateRoutes.Group("")
		restricted.Use(auth.AuthMiddleware("KOOR", "BPH", "ADMIN"))
		{
			restricted.GET("", pengurusHandler.GetAllPengurusByDivision)
			restricted.DELETE("/:id", pengurusHandler.DeletePengurus)
		}
	}

}
