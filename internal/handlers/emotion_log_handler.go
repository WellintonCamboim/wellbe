package handlers

import (
	"net/http"

	"github.com/WellintonCamboim/wellbe/internal/models"
	"github.com/WellintonCamboim/wellbe/internal/services"
	"github.com/labstack/echo/v4"
)

// EmotionLogHandler handles HTTP requests related to emotion logs
type EmotionLogHandler struct {
	emotionLogService *services.EmotionLogService
}

// NewEmotionLogHandler creates a new instance of EmotionLogHandler
func NewEmotionLogHandler(emotionLogService *services.EmotionLogService) *EmotionLogHandler {
	return &EmotionLogHandler{emotionLogService: emotionLogService}
}

// CreateEmotionLog godoc
// @Summary Creates a new emotion log
// @Description Registers a new emotion log in the system
// @Tags emotion_log
// @Accept json
// @Produce json
// @Param request body models.CreateEmotionLogRequest true "EmotionLog data"
// @Success 201 {object} models.EmotionLog "EmotionLog successfully created"
// @Failure 400 {object} map[string]string "Invalid data"
// @Failure 422 {object} map[string]string "Unprocessable entity"
// @Router /api/emotion-logs [post]
func (h *EmotionLogHandler) CreateEmotionLog(c echo.Context) error {
	var req models.CreateEmotionLogRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	emotionLog, err := h.emotionLogService.CreateEmotionLog(&req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, emotionLog)
}
