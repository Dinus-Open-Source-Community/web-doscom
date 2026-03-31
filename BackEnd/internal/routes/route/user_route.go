package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

// route for super admin only
func UserControllerRoute(r *gin.RouterGroup, UserHandler *handler.UserHandler) {
	user := r.Group("/user")
	// protected using middleware only superadmin can access
	admin := user.Group("/admin")
	admin.Use(auth.AuthMiddleware("ADMIN"))
	{
		// create new users
		admin.POST("/", UserHandler.SuperAdminCreateUser)
		// get users by id
		admin.GET("/:id", UserHandler.SuperAdmin   GetUser)
		// get all users
		admin.GET("/", UserHandler.SuperAdminGetAllUser)
		// update user by id
		admin.PUT("/:id", UserHandler.SuperAdminUpdateUser)
		// delete user by id
		admin.DELETE("/:id", UserHandler.SuperAdminDeleteUser)
		// create superadmin
		admin.POST("/superadmin", UserHandler.SuperAdminCreateSuperAdmin)
	}

	koor := user.Group("/koor")
	koor.Use(auth.AuthMiddleware("KOOR", "BPH"))
	{
		// all the routes here
		koor.POST("/", UserHandler.KoorCreateUser)      // create route
		koor.DELETE("/:id", UserHandler.KoorDeleteUser) // delete route
		koor.PUT("/:id", UserHandler.KoorUpdateUser)    // update route
		koor.GET("/", UserHandler.KoorGetAllUser)       // get all users
		koor.GET("/:id", UserHandler.KoorGetUser)       // get user by id
	}

}

