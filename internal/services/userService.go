package services

import (
	"go-layout/config"
	"go-layout/internal/models"
	"go-layout/pkg/token"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
}

type UserService struct {
	env        *config.Config
	db         *gorm.DB
	tokenMaker token.Maker
}

func NewUserService(env *config.Config, db *gorm.DB, tokenMaker token.Maker) UserServiceInterface {
	return &UserService{
		env:        env,
		db:         db,
		tokenMaker: tokenMaker,
	}
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := us.db.Find(&users)
	return users, result.Error
}
