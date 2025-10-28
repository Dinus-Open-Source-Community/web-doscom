package routes

import (
	"net/http"
	"web_doscom/internal/auth"
	"web_doscom/internal/config"
	"web_doscom/internal/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(app *config.Application) http.Handler {
	g := gin.Default()

	v1 := g.Group("api/v1")
	// Tambahkan route lain di sini
	{

		authHandler := auth.NewUserauth(&app.Model.Users)
		workHandler := handler.NewWorkHandler(app.DB)
		activitiesHandler := handler.NewActivitiesHandler(&app.Model.Activities)
		userHandler := handler.NewUserHandler(&app.Model.Users)

		AuthRoutes(v1, authHandler)

		UserControllerRoute(v1, userHandler)
		RegisterWorkRoutesV2(v1, workHandler)
		RegisterActivitiesRoutesV1(v1, activitiesHandler)

		// swagger
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return g
}
