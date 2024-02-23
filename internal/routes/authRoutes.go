package routes

import (
	"github.com/gin-gonic/gin"
	"go-layout/internal/app"
)

func AuthRoutes(router *gin.Engine, appCtx *app.Context) {
	router.Group("auth")
	{
		router.POST("/signup", appCtx.AuthController.SignUp)
		router.POST("/login", appCtx.AuthController.Login)
		router.POST("/logout", appCtx.AuthController.Logout)
		router.POST("/renew", appCtx.AuthController.Renew)
	}
}
