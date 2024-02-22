## Step 1: Define Service Interfaces

```go
// services/interfaces.go

package services

import "go-layout/token"

// UserServiceInterface defines the contract for user service operations.
type UserServiceInterface interface {
    SignUp(email, password string) error
    Login(email, password string) (string, error) // Returns token
    // Add other user-related methods here.
}

// userService implements UserServiceInterface.
type UserService struct {
    tokenMaker token.TokenMaker
}

func NewUserService(tokenMaker token.TokenMaker) UserServiceInterface {
    return &UserService{
        tokenMaker: tokenMaker,
    }
}

// Implement the interface methods.
func (s *UserService) SignUp(email, password string) error {
    // Implementation...
    return nil
}

func (s *UserService) Login(email, password string) (string, error) {
    // Implementation...
    return "", nil
}

```

## Step 2: Implement Controller Interfaces
```go
// controllers/interfaces.go

package controllers

import (
    "github.com/gin-gonic/gin"
    "go-layout/services"
)

// UserControllerInterface defines the contract for the user controller.
type UserControllerInterface interface {
    GetAllUsers(c *gin.Context)
    SignUp(c *gin.Context)
    Login(c *gin.Context)
    // Add other methods as needed.
}

// UserController implements UserControllerInterface.
type UserController struct {
    userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) UserControllerInterface {
    return &UserController{
        userService: userService,
    }
}

// Implement the interface methods.
func (uc *UserController) GetAllUsers(c *gin.Context) {
    // Implementation...
    c.JSON(200, gin.H{"message": "All users"})
}

func (uc *UserController) SignUp(c *gin.Context) {
    // Implementation...
    c.JSON(200, gin.H{"message": "Sign up successful"})
}

func (uc *UserController) Login(c *gin.Context) {
    // Implementation...
    c.JSON(200, gin.H{"message": "Login successful"})
}

```

### Decouple Business Logic from HTTP Handling
```go
// controllers/userController.go

package controllers

import (
    "github.com/gin-gonic/gin"
    "go-layout/services"
    "net/http"
)

type UserControllerInterface interface {
	GetAllUsers(c *gin.Context)
	SignUp(c *gin.Context)
	Login(c *gin.Context)
	// Add other methods as needed.
}

// Implementation of UserControllerInterface.
type UserController struct {
    userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
    return &UserController{
        userService: userService,
    }
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
    users, err := uc.userService.GetAllUsers() // Assuming such a method exists.
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) SignUp(c *gin.Context) {
    // Extract request details, validate, then call the service.
    var req struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }
    err := uc.userService.SignUp(req.Email, req.Password)
    if err != nil {
        // Handle error, e.g., user already exists
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign up"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Sign up successful"})
}

func (uc *UserController) Login(c *gin.Context) {
    // Similar to SignUp, extract details, validate, then call the service.
}

```


## Step 3: Application Setup Refactor
```go
// setup/appSetup.go

// setup/appSetup.go

// setup/appSetup.go

package setup

import (
	"go-layout/controllers"
	"go-layout/services"
	"go-layout/token"
)

type AppContext struct {
	UserController controllers.UserControllerInterface
	UserService    services.UserServiceInterface
	TokenMaker     token.TokenMaker
}

func InitializeApp(symmetricKey string) (*AppContext, error) {
	tokenMaker, err := token.NewPasetoMaker(symmetricKey)
	if err != nil {
		return nil, err
	}

	userService := services.NewUserService(tokenMaker)
	userController := controllers.NewUserController(userService)

	return &AppContext{
		UserController: userController,
		UserService:    userService,
		TokenMaker:     tokenMaker,
	}, nil
}


```


## Step 4: Adjust Main Function

```go
// main.go

package main

import (
    "github.com/gin-gonic/gin"
    "go-layout/setup"
    "log"
)

func main() {
    appCtx, err := setup.InitializeApp("your_symmetric_key_here")
    if err != nil {
        log.Fatalf("Failed to initialize application: %v", err)
    }

    router := gin.Default()
    routes.SetupRoutes(router, appCtx)
    router.Run() // Start the server
}

```

## Step 5a: Refactor the Middleware to Use tokenMaker

```go
// middleware/authMiddleware.go

package middleware

import (
    "github.com/gin-gonic/gin"
    "go-layout/token"
    "net/http"
)

func AuthMiddleware(tokenMaker token.TokenMaker) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extract the token from the request header
        authToken := c.GetHeader("Authorization")
        
        // Validate the token
        _, err := tokenMaker.VerifyToken(authToken)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
            return
        }

        // Token is valid; proceed with the request
        c.Next()
    }
}

```


## Step 5: Define a Route Configuration Function

```go
// routes/setup.go

// routes/setup.go

package routes

import (
	"github.com/gin-gonic/gin"
	"go-layout/middleware"
	"go-layout/setup"
)

func SetupRouter(appCtx *setup.AppContext) *gin.Engine {
	router := gin.Default()

	UserRoutes(router, appCtx)

	return router
}

```

### Applying Middleware to User Routes
```go
package routes

import (
	"github.com/gin-gonic/gin"
	"go-layout/setup"
)

func UserRoutes(router *gin.Engine, appCtx *setup.AppContext) {
	// User routes that don't require authentication
	router.POST("/users/signup", appCtx.UserController.SignUp)
	router.POST("/users/login", appCtx.UserController.Login)

	// Create a subgroup for routes that require authentication
	protectedUsers := router.Group("/users")
	// Apply the AuthMiddleware to this group
	protectedUsers.Use(middleware.AuthMiddleware(appCtx.TokenMaker))
	{
		protectedUsers.GET("/", appCtx.UserController.GetAllUsers)
		// Add more protected user routes here
	}
}

```



## Step 6: Update main.go to Use the New Route Configuration

```go
// main.go
// main.go

package main

import (
	"go-layout/routes"
	"go-layout/setup"
	"log"
)

func main() {
	appCtx, err := setup.InitializeApp("your_symmetric_key_here")
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Setup routes and middleware using AppContext
	router := routes.SetupRouter(appCtx)

	// Start the server
	if err := router.Run(); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}


```

