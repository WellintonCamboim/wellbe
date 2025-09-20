package handlers

import (
    "net/http"
    "github.com/WellintonCamboim/wellbe/internal/models"
    "github.com/WellintonCamboim/wellbe/internal/services"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
)

type TaskHandler struct {
    taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
    return &TaskHandler{taskService: taskService}
}

// CreateTask cria uma nova tarefa
func (h *TaskHandler) CreateTask(c echo.Context) error {
    var req models.CreateTaskRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
    }

    task, err := h.taskService.CreateTask(req)
    if err != nil {
        return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, task)
}

// GetTask busca uma tarefa pelo ID
func (h *TaskHandler) GetTask(c echo.Context) error {
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
    }

    task, err := h.taskService.GetTaskByID(id)
    if err != nil {
        if err.Error() == "task not found" {
            return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, task)
}

// ListTasksByUser lista todas as tarefas de um usuário
func (h *TaskHandler) ListTasksByUser(c echo.Context) error {
    userIDStr := c.QueryParam("user_id")
    userID, err := uuid.Parse(userIDStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user_id"})
    }

    tasks, err := h.taskService.ListTasksByUser(userID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, tasks)
}

// UpdateTask atualiza uma tarefa existente
func (h *TaskHandler) UpdateTask(c echo.Context) error {
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
    }

    var task models.Task
    if err := c.Bind(&task); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
    }
    task.ID = id

    if err := h.taskService.UpdateTask(&task); err != nil {
        return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, task)
}

// DeleteTask remove uma tarefa pelo ID
func (h *TaskHandler) DeleteTask(c echo.Context) error {
    idStr := c.Param("id")
    id, err := uuid.Parse(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
    }

    if err := h.taskService.DeleteTask(id); err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }

    return c.NoContent(http.StatusNoContent)
}