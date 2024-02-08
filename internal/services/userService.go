package services

import (
	"go-layout/internal/models"
	"go-layout/storage/db"
	"go-layout/utils"
)

type UserSignUpRequestDTO struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := db.AppDb.Find(&users)
	return users, result.Error
}

// SignUp adds a new user to the database
func SignUp(raw UserSignUpRequestDTO) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(raw.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		FirstName:    utils.PointerString(raw.FirstName),
		LastName:     utils.PointerString(raw.LastName),
		Email:        raw.Email,
		PasswordHash: hashedPassword,
	}
	err = db.AppDb.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
