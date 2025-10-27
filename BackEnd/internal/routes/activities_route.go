package routes

import (
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterActivitiesRoutesV1(rg *gin.RouterGroup, activitiesHandler *handler.ActivitiesHandler) {
	rg.POST("/activities/create", activitiesHandler.Create)
	rg.GET("/activities/list", activitiesHandler.List)
	rg.GET("/activities/get/:id", activitiesHandler.Get)
	rg.PUT("/activities/update/:id", activitiesHandler.Update)
	rg.DELETE("/activities/delete/:id", activitiesHandler.Delete)
}
