// @title Wellbe API
// @version 1.0
// @description API for user management
// @host localhost:8080
// @BasePath /api
package main

import (
	"github.com/WellintonCamboim/wellbe/api/docs"
	_ "github.com/WellintonCamboim/wellbe/api/docs"
	"github.com/WellintonCamboim/wellbe/internal/handlers"
	"github.com/WellintonCamboim/wellbe/internal/repositories"
	"github.com/WellintonCamboim/wellbe/internal/services"
	"github.com/WellintonCamboim/wellbe/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
    // Swagger configuration
    docs.SwaggerInfo.Title = "Wellbe API"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:8080"
    docs.SwaggerInfo.BasePath = "/api"

    db := database.Connect()
    e := echo.New()
    
    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    // e.Validator = NewCustomValidator() // You will need to implement this

    // Swagger UI
    e.GET("/swagger/*", echoSwagger.WrapHandler)

    // Dependency injection
    userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userHandler := handlers.NewUserHandler(userService)

    // Routes
    e.POST("/api/users", userHandler.CreateUser)

    e.Logger.Fatal(e.Start(":8080"))
}
