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
		lowerShared.GET("/me", userHandler.GetCurrentUser)
		lowerShared.PUT("/profile", userHandler.UpdateProfileUser)
		lowerShared.PUT("/change-password", userHandler.ChangePassword)
	}

	shared := user.Group("")
	shared.Use(auth.AuthMiddleware("ADMIN", "KOOR", "BPH"))
	{
		shared.PUT("/:id", userHandler.UpdateUserByID)
		shared.GET("/:id", userHandler.GetUserByID)
		shared.POST("", userHandler.CreateUser)
		shared.GET("", userHandler.GetAllUserBasedOnRole)
		shared.DELETE("/:id", userHandler.DeleteUser)
	}

	// admin-exclusive endpoint
	admin := r.Group("/admin/user")
	admin.Use(auth.AuthMiddleware("ADMIN"))
	{
		admin.POST("/super-admin", userHandler.CreateSuperAdmin)
		admin.GET("", userHandler.GetSuperAdmin)
		admin.PUT("/:id/change-password", userHandler.ChangePasswordAdmin)
	}

}
