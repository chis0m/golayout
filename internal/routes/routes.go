package routes

import (
	"github.com/gin-gonic/gin"
	"go-layout/internal/controllers"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})
	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Healthy",
		})
	})

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.GetAllUsers)
		userRoutes.POST("/signup", controllers.SignUp)
	}

	// Additional routes can be defined here
}
