package routes

import (
	"web_doscom/internal/auth"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func UserControllerRoute(r *gin.RouterGroup, UserHandler *handler.UserHandler) {
	admin := r.Group("/users")
	admin.Use(auth.AuthMiddleware()) // protected using middleware only superadmin can access
	{
		admin.POST("/", UserHandler.CreateUser)     // create new users
		admin.GET("/:id", UserHandler.GetUser)      // get users by id
		admin.GET("/", UserHandler.GetAllUser)      // get all users
		admin.PUT("/:id", UserHandler.UpdateUser)   // update user by id
		admin.DELETE(":id", UserHandler.DeleteUser) // delete user by id
	}

}
