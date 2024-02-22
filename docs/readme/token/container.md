Step 1: Create a Dependency Container
```go
// internal/app/container.go

package app

import (
	"go-layout/token"
	"go-layout/internal/services"
)

type Container struct {
	TokenService token.TokenMaker
	UserService  *services.UserService
}

func NewContainer(symmetricKey string) (*Container, error) {
	pasetoMaker, err := token.NewPasetoMaker(symmetricKey)
	if err != nil {
		return nil, err
	}

	userService := services.NewUserService(pasetoMaker)

	return &Container{
		TokenService: pasetoMaker,
		UserService:  userService,
	}, nil
}

```


Step 2: Use the Container in Your Application
```go
// main.go

package main

import (
	"go-layout/internal/app"
	"go-layout/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	container, err := app.NewContainer("your_symmetric_key_here")
	if err != nil {
		panic("failed to create the container: " + err.Error())
	}

	router := gin.Default()
	routes.SetupRoutes(router, container)
	router.Run() // specify your address or leave blank for default
}

```

Step 3: Adjusting Routes Setup to Use the Container
```go
// internal/routes/routes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"go-layout/internal/app"
)

func SetupRoutes(router *gin.Engine, container *app.Container) {
	// Example of setting up a route with access to the UserService
	router.GET("/users", func(c *gin.Context) {
		container.UserService.GetAllUsers(c)
	})

	// Other routes...
}

```

Step 4: Adjust Controllers to Receive Services Directly
```go
// internal/controllers/user_controller.go

package controllers

import (
	"go-layout/internal/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	// Implementation...
}

```