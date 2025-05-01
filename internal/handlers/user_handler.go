package handlers

import (
	"net/http"
	"strconv"

	"github.com/WellintonCamboim/wellbe/internal/models"
	"github.com/WellintonCamboim/wellbe/internal/services"
	"github.com/labstack/echo/v4"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
    userService *services.UserService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

// CreateUser godoc
// @Summary Creates a new user
// @Description Registers a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param request body models.CreateUserRequest true "User data"
// @Success 201 {object} models.User "User successfully created"
// @Failure 400 {object} map[string]string "Invalid data"
// @Failure 422 {object} map[string]string "Unprocessable entity"
// @Router /api/users [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
    var req models.CreateUserRequest
    
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
    }

    user, err := h.userService.CreateUser(req)
    if err != nil {
        return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, user)
}

// GetUser godoc
// @Summary Get a user by ID
// @Description Get user details by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User "User found"
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/users/{id} [get]
func (h *UserHandler) GetUser(c echo.Context) error {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
    }

    user, err := h.userService.GetUserByID(uint(id))
    if err != nil {
        if err.Error() == "user not found" {
            return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, user)
}