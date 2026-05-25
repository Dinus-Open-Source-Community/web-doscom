package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserControllerRoute(r *gin.RouterGroup, userHandler *handler.UserHandler) {
	user := r.Group("/user")

	// shared endpoint dengan role-based
	lowerShared := user.Group("")
	lowerShared.Use(auth.AuthMiddleware("ADMIN", "KOOR", "BPH", "PENGURUS"))
	{
		lowerShared.GET("/:id", userHandler.GetUser)
		lowerShared.PUT("/change-password", userHandler.ChangePassword)
	}
	shared := user.Group("")
	shared.Use(auth.AuthMiddleware("ADMIN", "KOOR", "BPH"))
	{
		shared.POST("", userHandler.CreateUser)
		shared.GET("", userHandler.GetAllUserBasedOnRole)
		shared.PUT("/:id", userHandler.UpdateUser)
		shared.DELETE("/:id", userHandler.DeleteUser)
	}

	// admin-exclusive endpoint
	admin := r.Group("/admin/user")
	admin.Use(auth.AuthMiddleware("ADMIN"))
	{
		admin.POST("/super-admin", userHandler.CreateSuperAdmin)
		admin.PUT("/:id/change-password", userHandler.ChangePasswordAdmin)
	}

}
