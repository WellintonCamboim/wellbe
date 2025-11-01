package handlers

import (
	"net/http"
	"github.com/WellintonCamboim/wellbe/internal/models"
	"github.com/WellintonCamboim/wellbe/internal/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SkillHandler struct {
	skillService *services.SkillService
}

func NewSkillHandler(skillService *services.SkillService) *SkillHandler {
	return &SkillHandler{skillService: skillService}
}

// CreateSkill cria uma nova skill
func (h *SkillHandler) CreateSkill(c echo.Context) error {
	var req models.CreateSkillRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	skill, err := h.skillService.CreateSkill(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, skill)
}

// GetSkill busca uma skill pelo ID
func (h *SkillHandler) GetSkill(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	skill, err := h.skillService.GetSkillByID(id)
	if err != nil {
		if err.Error() == "skill not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, skill)
}

// ListSkillsByUser lista todas as skills de um usuário
func (h *SkillHandler) ListSkillsByUser(c echo.Context) error {
	userIDStr := c.QueryParam("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user_id"})
	}

	skills, err := h.skillService.ListByUser(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, skills)
}

// UpdateSkill atualiza uma skill existente
func (h *SkillHandler) UpdateSkill(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	var req models.UpdateSkillRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid data"})
	}

	updatedSkill, err := h.skillService.UpdateSkill(id, req)
	if err != nil {
		if err.Error() == "skill not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	
	return c.JSON(http.StatusOK, updatedSkill)
}

// DeleteSkill remove uma skill existente
func (h *SkillHandler) DeleteSkill(c echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	if err := h.skillService.DeleteSkill(id); err != nil {
		if err.Error() == "skill not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}	