package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserControllerRoute(r *gin.RouterGroup, UserHandler *handler.UserHandler) {
	user := r.Group("/user")

	// shared endpoint dengan role-based
	shared := user.Group("")
	shared.Use(auth.AuthMiddleware("ADMIN", "KOOR", "BPH"))
	{
		shared.POST("", UserHandler.CreateUser)
		shared.GET("/:id", UserHandler.GetUser)
		shared.GET("", UserHandler.GetAllUserBasedOnRole)
		shared.PUT("/:id", UserHandler.UpdateUser)
		shared.DELETE("/:id", UserHandler.DeleteUser)
	}

	// admin-exclusive endpoint
	admin := user.Group("/admin")
	admin.Use(auth.AuthMiddleware("ADMIN"))
	{
		admin.POST("", UserHandler.CreateSuperAdmin)
	}

}
