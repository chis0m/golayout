1. Define a Global Container or Context

```go
// appContext.go (create this file in your project)

package app

type AppContext struct {
	TokenMaker token.TokenMaker
	// Add other services and dependencies here
	UserService *services.UserService
	// OtherService ...
}

```

2. Initialize Your Application Context in Main

```go

func main() {
    symmetricKey := "your_symmetric_key_here" // Load this securely
    tokenMaker, err := token.NewPasetoMaker(symmetricKey)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create token maker")
    }

    appContext := &app.AppContext{
        TokenMaker: tokenMaker,
        // Initialize other services and pass dependencies as needed
        UserService: services.NewUserService(tokenMaker),
        // OtherService: ...
    }

    setupRouter(appContext) // Pass the context to your router setup
}

func setupRouter(appCtx *app.AppContext) {
    router := gin.Default()
    routes.SetupRoutes(router, appCtx) // Modify SetupRoutes to accept appCtx
    // ...
}

```

3. Modify Routes to Use Application Context

```go
// routes/routes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"go-layout/app" // Adjust the import path based on your project structure
)

func SetupRoutes(router *gin.Engine, appCtx *app.AppContext) {
    // Now you can pass appCtx or specific services to your controllers
    userRoutes := router.Group("/users")
    {
        userRoutes.GET("/", func(c *gin.Context) { controllers.GetAllUsers(c, appCtx) })
        userRoutes.POST("/signup", func(c *gin.Context) { controllers.SignUp(c, appCtx) })
        userRoutes.POST("/login", func(c *gin.Context) { controllers.Login(c, appCtx) })
    }

    // Setup other routes...
}

```

4. Adjust Controllers to Receive the Application Context
```go
// controllers/userController.go

package controllers

import (
	"github.com/gin-gonic/gin"
	"go-layout/app" // Adjust the import path based on your project structure
)

func SignUp(c *gin.Context, appCtx *app.AppContext) {
    // Use appCtx.TokenMaker or appCtx.UserService as needed
    var user services.UserSignUpRequestDTO
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Example of using UserService from appCtx
    res, err := appCtx.UserService.SignUp(user)
    // Handle the response...
}

```