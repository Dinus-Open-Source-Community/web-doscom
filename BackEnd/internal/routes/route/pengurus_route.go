package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPengurusRoutes(rg *gin.RouterGroup, pengurusHandler *handler.PengurusHandler) {
	// public api
	publicRoutes := rg.Group("/pengurus")
	{
<<<<<<< HEAD
		publicRoutes.GET("/division/:division", pengurusHandler.GetAllPengurus)
=======
		p.GET("/:id", pengurusHandler.GetPengurus)
		p.GET("/", pengurusHandler.GetAllPengurus)
>>>>>>> master
	}

	// private api -> need auth
	authRoutes := rg.Group("/admin/pengurus")
	authRoutes.Use(auth.AuthMiddleware("ANGGOTA", "KOOR", "BPH", "ADMIN"))
	{
<<<<<<< HEAD
		authRoutes.GET("/:id", pengurusHandler.GetPengurusByID)
		authRoutes.POST("/", pengurusHandler.CreatePengurus)
		authRoutes.PUT("/:id", pengurusHandler.UpdatePengurus)
=======
		pengurusAuth.POST("/", pengurusHandler.CreatePengurus)
		pengurusAuth.PUT("/:id", pengurusHandler.UpdatePengurus)
	}
>>>>>>> master

		authRoutes.DELETE("/:id",
			auth.AuthMiddleware("KOOR", "BPH", "ADMIN"),
			pengurusHandler.DeletePengurus,
		)
		authRoutes.GET("",
			auth.AuthMiddleware("KOOR", "BPH", "ADMIN"),
			pengurusHandler.GetAllPengurusByDivision,
		)
	}

}
