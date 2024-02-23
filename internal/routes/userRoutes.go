package routes

import (
	"github.com/gin-gonic/gin"
	"go-layout/internal/app"
	"go-layout/internal/middleware"
)

func UserRoutes(router *gin.Engine, appCtx *app.Context) {
	router.Group("/users").Use(middleware.AuthMiddleware(appCtx.TokenMaker))
	{
		router.GET("/", appCtx.UserController.GetAllUsers)
	}
}
