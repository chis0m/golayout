package app

import (
	"go-layout/config"
	"go-layout/internal/controllers"
	"go-layout/internal/services"
	"go-layout/pkg/token"
	"gorm.io/gorm"
)

type Context struct {
	env            *config.Config
	Db             *gorm.DB
	TokenMaker     token.Maker
	AuthService    services.AuthServiceInterface
	AuthController controllers.AuthControllerInterface
	UserService    services.UserServiceInterface
	UserController controllers.UserControllerInterface
}

func Initialize(env *config.Config, db *gorm.DB, symmetricKey string) (*Context, error) {
	tokenMaker, err := token.NewPasetoMaker(symmetricKey)
	if err != nil {
		return nil, err
	}

	authService := services.NewAuthService(env, db, tokenMaker)
	authController := controllers.NewAuthController(authService)
	userService := services.NewUserService(env, db, tokenMaker)
	userController := controllers.NewUserController(userService)

	return &Context{
		env:            env,
		Db:             db,
		TokenMaker:     tokenMaker,
		AuthService:    authService,
		AuthController: authController,
		UserService:    userService,
		UserController: userController,
	}, nil
}
