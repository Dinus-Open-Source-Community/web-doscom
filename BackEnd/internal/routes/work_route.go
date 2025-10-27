package routes

import (
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterWorkRoutesV2(rg *gin.RouterGroup, workHandler *handler.WorkHandler) {

	rg.POST("/work/create", workHandler.Create)       // POST /api/v1/work/create
	rg.GET("/work/list", workHandler.List)            // GET /api/v1/work/list
	rg.GET("/work/get/:id", workHandler.Get)          // GET /api/v1/work/get/:id
	rg.PUT("/work/update/:id", workHandler.Update)    // PUT /api/v1/work/update/:id
	rg.DELETE("/work/delete/:id", workHandler.Delete) // DELETE /api/v1/work/delete/:id

}
