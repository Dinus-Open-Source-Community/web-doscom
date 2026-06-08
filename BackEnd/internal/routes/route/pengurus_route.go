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

	privateRoutes := rg.Group("/pengurus")
	privateRoutes.Use(auth.AuthMiddleware("PENGURUS", "KOOR", "BPH", "ADMIN"))
	{
		privateRoutes.POST("", pengurusHandler.CreateMyPengurusProfile)
		privateRoutes.GET("/profile", pengurusHandler.GetPengurusProfile)
		privateRoutes.PUT("/me", pengurusHandler.UpdateMyPengurusProfile)
	}

	managedRoutes := rg.Group("/admin/pengurus")
	managedRoutes.Use(auth.AuthMiddleware("KOOR", "BPH", "ADMIN"))
	{
		managedRoutes.POST("", pengurusHandler.CreatePengurusProfile)
		managedRoutes.GET("/:id", pengurusHandler.GetPengurusByID)
		managedRoutes.GET("/by-user/:user_id", pengurusHandler.GetPengurusByUserID)
		managedRoutes.PUT("/:id", pengurusHandler.UpdatePengurusByID)
		managedRoutes.DELETE("/delete/:id", pengurusHandler.DeletePengurus)
	}

}
