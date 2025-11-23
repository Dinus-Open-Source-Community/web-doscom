package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

// route for super admin only
func UserControllerRoute(r *gin.RouterGroup, UserHandler *handler.UserHandler) {
	admin := r.Group("/admin")
	admin.Use(auth.AuthMiddleware("ADMIN")) // protected using middleware only superadmin can access
	{
		admin.POST("/", UserHandler.SuperAdminCreateUser)                 // create new users
		admin.GET("/:id", UserHandler.SuperAdminGetUser)                  // get users by id
		admin.GET("/", UserHandler.SuperAdminGetAllUser)                  // get all users
		admin.PUT("/:id", UserHandler.SuperAdminUpdateUser)               // update user by id
		admin.DELETE("/:id", UserHandler.SuperAdminDeleteUser)            // delete user by id
		admin.POST("/superadmin", UserHandler.SuperAdminCreateSuperAdmin) // create superadmin
	}

}

// route for koor only
func UserControllerKoor(r *gin.RouterGroup, UserHandler *handler.UserHandler) {
	koor := r.Group("/koor")
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
