package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-layout/internal/app"
	httpLogger "go-layout/pkg/logger"
)

func SetupRoutes(appCtx *app.Context) *gin.Engine {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(corsConfig))
	router.Use(httpLogger.CustomLog())

	// routes
	AuthRoutes(router, appCtx)
	UserRoutes(router, appCtx)

	return router
}
