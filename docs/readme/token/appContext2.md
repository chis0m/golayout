Step 1: Service Initializers

```go
// services/init.go

package services

func NewUserService(tokenMaker token.TokenMaker) *UserService {
    // Initialize and return a new instance of UserService
    return &UserService{
        tokenMaker: tokenMaker,
    }
}

// Repeat for other services...

```

Step 2: Controller Initializers
```go
// controllers/init.go

package controllers

import "go-layout/internal/services"

func NewUserController(userService *services.UserService) *UserController {
    return &UserController{
        userService: userService,
    }
}

// Repeat for other controllers...

```

Step 3: Wiring Dependencies in a Centralized Setup Function
```go
// setup/appSetup.go

package setup

import (
    "go-layout/controllers"
    "go-layout/services"
    "go-layout/token"
)

type AppContext struct {
    UserController *controllers.UserController
    // Add other controllers here
}

func InitializeApp(symmetricKey string) (*AppContext, error) {
    tokenMaker, err := token.NewPasetoMaker(symmetricKey)
    if err != nil {
        return nil, err
    }

    // Initialize services
    userService := services.NewUserService(tokenMaker)

    // Initialize controllers with their dependencies
    userController := controllers.NewUserController(userService)

    // Return a struct containing all initialized controllers
    return &AppContext{
        UserController: userController,
        // Other controllers...
    }, nil
}

```


Step 4: Use in Main

```go
// main.go

package main

import (
    "go-layout/setup"
    "log"
)

func main() {
    appCtx, err := setup.InitializeApp("your_symmetric_key_here")
    if err != nil {
        log.Fatalf("Failed to initialize application: %v", err)
    }

    // Now, appCtx contains all your controllers, fully initialized.
    // Pass appCtx to your router setup or wherever needed.
}

```
Step 5: Adjust the Route Setup Function

```go
// routes/routes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"go-layout/setup" // Adjust import path as necessary
)

func SetupRoutes(router *gin.Engine, appCtx *setup.AppContext) {
    // Use appCtx to access controllers for route handlers
    router.GET("/users", func(c *gin.Context) {
        appCtx.UserController.GetAllUsers(c)
    })
    router.POST("/users/signup", func(c *gin.Context) {
        appCtx.UserController.SignUp(c)
    })
    router.POST("/users/login", func(c *gin.Context) {
        appCtx.UserController.Login(c)
    })

    // Define other routes...
}

```

5b: Middleware
```go
// routes/setupRoutes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"go-layout/middleware"
	"go-layout/setup"
)

func SetupRoutes(router *gin.Engine, appCtx *setup.AppContext) {
	// Apply the AuthMiddleware to all routes needing authentication
	authMiddleware := middleware.AuthMiddleware(appCtx)
	apiRoutes := router.Group("/api")
	apiRoutes.Use(authMiddleware)

	apiRoutes.GET("/protected", func(c *gin.Context) {
		// Example protected route
		c.JSON(http.StatusOK, gin.H{"message": "You are authenticated"})
	})

	// Setup other routes...
}

```

Step 6: Controller Methods as HTTP Handlers

```go
// controllers/userController.go

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
    // Dependencies, e.g., services
    UserService *services.UserService
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
    // Use uc.UserService to handle the request...
    c.JSON(http.StatusOK, gin.H{"message": "All users"})
}

func (uc *UserController) SignUp(c *gin.Context) {
    // Handle sign-up...
    c.JSON(http.StatusOK, gin.H{"message": "Sign up successful"})
}

func (uc *UserController) Login(c *gin.Context) {
    // Handle login...
    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

```

Step 8: Main Function Usage

```go
// main.go

func main() {
    appCtx, err := setup.InitializeApp("your_symmetric_key")
    if err != nil {
        log.Fatal("Failed to initialize app context:", err)
    }

    router := gin.Default()
    routes.SetupRoutes(router, appCtx) // Pass the fully initialized AppContext
    router.Run() // Start the server
}

```
