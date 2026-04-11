package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterWorkRoutes(rg *gin.RouterGroup, workHandler *handler.WorkHandler) {

	workRoutes := rg.Group("/works")
	workRoutes.Use(auth.AuthMiddleware("SuperAdmin", "KorPemro", "KorJaringan", "KorData", "KorMedcrev", "BPH"))
	{
		workRoutes.POST("", workHandler.Create)
		workRoutes.GET("/:id", workHandler.Get)
		workRoutes.GET("", workHandler.List)
		workRoutes.PUT("/:id", workHandler.Update)
		workRoutes.DELETE("/:id", workHandler.Delete)
	}

}
