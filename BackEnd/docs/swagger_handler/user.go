package swagger_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"web_doscom/internal/handler"
)

var handlerInstance *handler.UserHandler

func SetUserHandler(r *handler.UserHandler) {
	handlerInstance = r
}

// super admin wrapper

func SuperAdminCreatUser(c *gin.Context) {
	handlerInstance.CreateUser(c)
}
