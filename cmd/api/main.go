package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/WellintonCamboim/wellbe/internal/handlers"
    "github.com/WellintonCamboim/wellbe/internal/repositories"
    "github.com/WellintonCamboim/wellbe/internal/services"
    "github.com/WellintonCamboim/wellbe/pkg/database"
)

func main() {
    db := database.Connect()
    e := echo.New()
    
    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    // e.Validator = NewCustomValidator() // Você precisará implementar isso

    // Injeção de dependências
    userRepo := repositories.NewUserRepository(db)
    userService := services.NewUserService(userRepo)
    userHandler := handlers.NewUserHandler(userService)

    // Rotas
    e.POST("/api/users", userHandler.CreateUser)

    e.Logger.Fatal(e.Start(":8080"))
}