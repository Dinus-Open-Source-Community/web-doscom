package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterWorkRoutes(rg *gin.RouterGroup, workHandler *handler.WorkHandler) {

	workRoutes := rg.Group("/works")
	workRoutes.Use(auth.AuthMiddleware("Super_Admin", "Kor_Pemro", "Kor_Jaringan", "Kor_Data", "Kor_Medcrev", "BPH"))
	{
		workRoutes.POST("", workHandler.Create)
		workRoutes.GET("/:id", workHandler.Get)
		workRoutes.GET("", workHandler.List)
		workRoutes.PUT("/:id", workHandler.Update)
		workRoutes.DELETE("/:id", workHandler.Delete)
	}

}
