// @title Wellbe API
// @version 1.0
// @description API for user management
// @host localhost:8080
// @BasePath /api
package main

import (
	"os"
	"github.com/WellintonCamboim/wellbe/api/docs"
	_ "github.com/WellintonCamboim/wellbe/api/docs"
	"github.com/WellintonCamboim/wellbe/internal/handlers"
	"github.com/WellintonCamboim/wellbe/internal/repositories"
	"github.com/WellintonCamboim/wellbe/internal/services"
	"github.com/WellintonCamboim/wellbe/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/joho/godotenv"
)

func main() {
	if env := os.Getenv("APP_ENV"); env == "" || env == "development" {
		_ = godotenv.Load(".env")
	}

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

	// Dependency injection para emotion_log
	emotionLogRepo := repositories.NewEmotionLogRepository(db)
	emotionLogService := services.NewEmotionLogService(emotionLogRepo)
	emotionLogHandler := handlers.NewEmotionLogHandler(emotionLogService)

	// Dependency injection para skill
	skillRepo := repositories.NewSkillRepository(db)
	skillService := services.NewSkillService(skillRepo)
	skillHandler := handlers.NewSkillHandler(skillService)

	// Dependency injection para task
	taskRepo := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Rotas de usuário
	e.POST("/api/users", userHandler.CreateUser)
	e.GET("/api/users/:id", userHandler.GetUser)

	// Rotas de emotion_log
	e.POST("/api/emotion-logs", emotionLogHandler.CreateEmotionLog)

	// Rotas de tarefas (task)
	e.POST("/api/tasks", taskHandler.CreateTask)
	e.GET("/api/tasks/:id", taskHandler.GetTask)
	e.GET("/api/tasks", taskHandler.ListTasksByUser)
	e.PUT("/api/tasks/:id", taskHandler.UpdateTask)
	e.DELETE("/api/tasks/:id", taskHandler.DeleteTask)

	// Rotas de skills
	e.POST("/api/skills", skillHandler.CreateSkill)
	e.GET("/api/skills/:id", skillHandler.GetSkill)
	e.GET("/api/skills/user", skillHandler.ListSkillsByUser)
	e.PUT("/api/skills/:id", skillHandler.UpdateSkill)
	e.DELETE("/api/skills/:id", skillHandler.DeleteSkill)

	e.Logger.Fatal(e.Start(":8080"))
}
