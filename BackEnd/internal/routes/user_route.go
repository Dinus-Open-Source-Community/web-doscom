package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

// route for super admin only
func UserControllerRoute(r *gin.RouterGroup, UserHandler *handler.UserHandler) {
	admin := r.Group("/admin")
	admin.Use(auth.AuthMiddleware("Super_Admin")) // protected using middleware only superadmin can access
	{
		admin.POST("/", UserHandler.CreateUser)     // create new users
		admin.GET("/:id", UserHandler.GetUser)      // get users by id
		admin.GET("/", UserHandler.GetAllUser)      // get all users
		admin.PUT("/:id", UserHandler.UpdateUser)   // update user by id
		admin.DELETE(":id", UserHandler.DeleteUser) // delete user by id
	}

}

// route for koor only
func UserControllerKoor(r *gin.RouterGroup, UserHandler *handler.UserHandler) {
	koor := r.Group("/koor")
	koor.Use(auth.AuthMiddleware("Kor_Pemro", "Kor_Jaringan", "Kor_Data", "Kor_Medcrev", "BPH"))
	{
		// all the routes here
		// create route
		koor.POST("/", UserHandler.CreateUser)
		// delete route
		// update route
	}
}
