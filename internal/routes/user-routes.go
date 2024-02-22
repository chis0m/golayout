package routes

import (
	"github.com/gin-gonic/gin"
	"go-layout/internal/app"
	"go-layout/internal/middleware"
)

func UserRoutes(router *gin.Engine, appCtx *app.Context) {
	router.POST("/users/signup", appCtx.UserController.SignUp)
	router.POST("/users/login", appCtx.UserController.Login)
	router.Group("/users").Use(middleware.AuthMiddleware(appCtx.TokenMaker))
	{
		router.GET("/", appCtx.UserController.GetAllUsers)
	}
}
