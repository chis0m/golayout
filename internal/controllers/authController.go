package controllers

import (
	"github.com/gin-gonic/gin"
	"go-layout/internal/appx"
	"go-layout/internal/services"
	"go-layout/internal/types"
	"net/http"
)

type AuthControllerInterface interface {
	SignUp(c *gin.Context)
	Login(c *gin.Context)
	Renew(c *gin.Context)
	Logout(c *gin.Context)
}

type AuthController struct {
	authService services.AuthServiceInterface
}

func NewAuthController(authService services.AuthServiceInterface) AuthControllerInterface {
	return &AuthController{
		authService: authService,
	}
}

func (ac *AuthController) SignUp(c *gin.Context) {
	var user types.UserSignUpRequestDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, appx.ErrorResponse("Invalid Request", err.Error()))
		return
	}

	res, err := ac.authService.SignUp(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appx.ErrorResponse("Signup Failed", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, appx.SuccessResponse("Signup Successful", res))
}

func (ac *AuthController) Login(c *gin.Context) {
	var login types.UserLoginRequestDTO
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, appx.ErrorResponse("Invalid Request", err.Error()))
		return
	}
	res, err := ac.authService.Login(c.Request.UserAgent(), c.ClientIP(), login)
	if err != nil {
		c.JSON(http.StatusBadRequest, appx.ErrorResponse("Login Failed", err.Error()))
	}
	c.JSON(http.StatusOK, appx.SuccessResponse("Login Successful", res))
}

func (ac *AuthController) Renew(c *gin.Context) {
	var renew types.RefreshTokenRequestDTO
	if err := c.ShouldBindJSON(&renew); err != nil {
		c.JSON(http.StatusBadRequest, appx.ErrorResponse("Invalid Request", err.Error()))
		return
	}
	res, err := ac.authService.Renew(renew)
	if err != nil {
		c.JSON(http.StatusBadRequest, appx.ErrorResponse("Token Renewal Failed", err.Error()))
	}
	c.JSON(http.StatusOK, appx.SuccessResponse("Token Renew Successful", res))
}

func (ac *AuthController) Logout(c *gin.Context) {
	var logout types.RefreshTokenRequestDTO
	if err := c.ShouldBindJSON(&logout); err != nil {
		c.JSON(http.StatusBadRequest, appx.ErrorResponse("Invalid Request", err.Error()))
		return
	}
	err := ac.authService.Logout(logout)
	if err != nil {
		c.JSON(http.StatusBadRequest, appx.ErrorResponse("Logout Failed", err.Error()))
	}
	c.JSON(http.StatusOK, appx.SuccessResponse("Logout Successful", ""))
}
