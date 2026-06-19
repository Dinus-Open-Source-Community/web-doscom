package routes

import (
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterHistoryRoutes(rg *gin.RouterGroup, historyHandler *handler.HistoryHandler) {
	publickRoutes := rg.Group("/history")
	{
		publickRoutes.GET("/:id", historyHandler.GetHistoryTimelineByID)
		publickRoutes.GET("", historyHandler.GetAllHistoryTimeline)
	}
}
