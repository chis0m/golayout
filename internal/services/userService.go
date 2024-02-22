package services

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"go-layout/internal/models"
	"go-layout/pkg/token"
	"go-layout/types"
	"go-layout/utils"
	"gorm.io/gorm"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	SignUp(raw types.UserSignUpRequestDTO) (*models.User, error)
	Login(raw types.UserLoginRequestDTO) (*models.User, error)
}

type UserService struct {
	db         *gorm.DB
	tokenMaker token.Maker
}

func NewUserService(db *gorm.DB, tokenMaker token.Maker) UserServiceInterface {
	return &UserService{
		db:         db,
		tokenMaker: tokenMaker,
	}
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := us.db.Find(&users)
	return users, result.Error
}

func (us *UserService) SignUp(raw types.UserSignUpRequestDTO) (*models.User, error) {
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
	err = us.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) Login(raw types.UserLoginRequestDTO) (*models.User, error) {
	var user models.User
	err := us.db.Where("email = ?", raw.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		} else {
			return nil, fmt.Errorf("could not fetch user %+v", err)
		}
	}

	err = utils.VerifyPassword(user.PasswordHash, raw.Password)
	if err != nil {
		log.Error().Msgf("Password verifcation failed %s", err)
		return nil, fmt.Errorf("invlid credentials %s", err)
	}
	return nil, nil
}
