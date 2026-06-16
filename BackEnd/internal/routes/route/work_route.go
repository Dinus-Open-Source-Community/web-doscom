package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterWorkRoutes(rg *gin.RouterGroup, workHandler *handler.WorkHandler) {

	// public api
	publickRoutes := rg.Group("/works")
	{
		publickRoutes.GET("", workHandler.GetAllWorks)
		publickRoutes.GET("/:id", workHandler.GetWorkByIDPublic)
		publickRoutes.GET("/types", workHandler.GetWorkTypes)
	}

	// private api
	adminWorkRoutes := rg.Group("/admin/works")
	adminWorkRoutes.Use(auth.AuthMiddleware("ADMIN", "KOOR"))
	{
		adminWorkRoutes.POST("", workHandler.CreateWork)
		adminWorkRoutes.GET("/:id", workHandler.GetWorkByIDAdmin)
		adminWorkRoutes.GET("", workHandler.GetAllWorksByDivision)
		adminWorkRoutes.PUT("/:id", workHandler.UpdateWork)
		adminWorkRoutes.DELETE("/:id", workHandler.Delete)
	}

	moderationWorkRoutes := rg.Group("/admin/works")
	moderationWorkRoutes.Use(auth.AuthMiddleware("ADMIN", "BPH"))
	{
		moderationWorkRoutes.PUT("/:id/status", workHandler.UpdateStatusWork)
	}
}
