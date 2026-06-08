package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/constants"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterWorkRoutes(rg *gin.RouterGroup, workHandler *handler.WorkHandler) {

	// public api
	publickRoutes := rg.Group("/works")
	{
		publickRoutes.GET("", workHandler.CreateWork)
		publickRoutes.GET(
			"/:projecttype", workHandler.GetAllWorks,
		)
	}

	// private api
	privateRoutes := rg.Group("/admin/works")
	privateRoutes.Use(auth.AuthMiddleware(constants.RoleKeySuperAdmin, constants.RoleKoordinator))
	{
		privateRoutes.POST("", workHandler.CreateWork)
		privateRoutes.GET("/:id", workHandler.GetWorkByID)
		privateRoutes.GET("", workHandler.GetAllWorksByDivision)
		privateRoutes.PUT("/:id", workHandler.UpdateWork)
		privateRoutes.DELETE("/:id", workHandler.Delete)
	}

}
