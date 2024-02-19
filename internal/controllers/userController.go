package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go-layout/internal/response"
	"go-layout/internal/services"
	"net/http"
)

// GetAllUsers fetches all users
func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		log.Err(err).Msg("GetAllUsers: failed to fetch users from users table")
		c.JSON(http.StatusInternalServerError, response.Error("internal server error", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success("fetch all users successful", users))
}

// SignUp a new user
func SignUp(c *gin.Context) {
	var user services.UserSignUpRequestDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := services.SignUp(user)
	if err != nil {
		log.Err(err).Msg("SignUp: failed to signup user")
		c.JSON(http.StatusInternalServerError, response.Error("could not signup user", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, response.Success("signup successful", res))
}
