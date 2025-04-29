package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/WellintonCamboim/wellbe/internal/models"
    "github.com/WellintonCamboim/wellbe/internal/services"
)

type UserHandler struct {
    userService *services.UserService  // Mude para ponteiro
}

func NewUserHandler(userService *services.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
    var req models.CreateUserRequest
    
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "dados inválidos"})
    }

    user, err := h.userService.CreateUser(req)
    if err != nil {
        return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, user)
}