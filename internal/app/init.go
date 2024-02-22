package app

import (
	"go-layout/internal/controllers"
	"go-layout/internal/services"
	"go-layout/pkg/token"
	"gorm.io/gorm"
)

type Context struct {
	Db             *gorm.DB
	TokenMaker     token.Maker
	UserController controllers.UserControllerInterface
	UserService    services.UserServiceInterface
}

func Initialize(db *gorm.DB, symmetricKey string) (*Context, error) {
	tokenMaker, err := token.NewPasetoMaker(symmetricKey)
	if err != nil {
		return nil, err
	}

	userService := services.NewUserService(db, tokenMaker)
	userController := controllers.NewUserController(userService)

	return &Context{
		Db:             db,
		TokenMaker:     tokenMaker,
		UserController: userController,
		UserService:    userService,
	}, nil
}
