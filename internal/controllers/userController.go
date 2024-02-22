package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go-layout/internal/response"
	"go-layout/internal/services"
	"go-layout/types"
	"net/http"
)

type UserControllerInterface interface {
	GetAllUsers(c *gin.Context)
	SignUp(c *gin.Context)
	Login(c *gin.Context)
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
		log.Err(err).Msg("GetAllUsers: failed to fetch users from users table")
		c.JSON(http.StatusInternalServerError, response.Error("internal server error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success("fetch all users successful", users))
}

func (uc *UserController) SignUp(c *gin.Context) {
	var user types.UserSignUpRequestDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := uc.userService.SignUp(user)
	if err != nil {
		log.Err(err).Msg("SignUp: failed to signup user")
		c.JSON(http.StatusInternalServerError, response.Error("could not signup user", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.Success("signup successful", res))
}

func (uc *UserController) Login(c *gin.Context) {
	var login types.UserLoginRequestDTO
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//res, err := uc.userService.Login(login)

}
