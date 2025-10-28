package routes

import (
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterActivitiesRoutesV1(rg *gin.RouterGroup, activitiesHandler *handler.ActivitiesHandler) {
	activities := rg.Group("/activities")
	{
		activities.POST("", activitiesHandler.CreateActivities)
		activities.GET("/:id", activitiesHandler.GetActivities)
		activities.PUT("/:id", activitiesHandler.UpdateActivities)
		activities.DELETE("/:id", activitiesHandler.DeleteActivities)
	}
}
