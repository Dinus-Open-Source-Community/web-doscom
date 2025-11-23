package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPengurusRoutes(rg *gin.RouterGroup, pengurusHandler *handler.PengurusHandler) {
	// public api
	p := rg.Group("/pengurus")
	// public api
	{
<<<<<<< HEAD
=======
		p.GET("/:id", pengurusHandler.GetPengurus)
>>>>>>> fd36652 (add some function, fix some function, malaasss)
		p.GET("/", pengurusHandler.GetAllPengurus)
	}

	// private api -> need auth
	pengurusAuth := p.Group("/")
	pengurusAuth.Use(auth.AuthMiddleware("ANGGOTA", "KOOR", "BPH", "ADMIN"))
	{
<<<<<<< HEAD
		p.GET("/:id", pengurusHandler.GetPengurus)
=======
>>>>>>> fd36652 (add some function, fix some function, malaasss)
		pengurusAuth.POST("/", pengurusHandler.CreatePengurus)
		pengurusAuth.PUT("/:id", pengurusHandler.UpdatePengurus)
	}

	// private api -> koor khusus
	koor := p.Group("/")
	koor.Use(auth.AuthMiddleware("KOOR", "BPH", "ADMIN"))
	{
		pengurusAuth.DELETE("/:id", pengurusHandler.DeletePengurus)
	}

}
