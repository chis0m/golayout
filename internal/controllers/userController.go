package controllers

import (
	"github.com/gin-gonic/gin"
	"go-layout/internal/httpx"
	"go-layout/internal/services"
	"net/http"
)

type UserControllerInterface interface {
	GetAllUsers(c *gin.Context)
}

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) UserControllerInterface {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpx.ErrorResponse(
			"internal server error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, httpx.SuccessResponse("fetch all users successful", users))
}
