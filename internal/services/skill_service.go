package services

import (
	"errors"
	"time"

	"github.com/WellintonCamboim/wellbe/internal/models"
	"github.com/WellintonCamboim/wellbe/internal/repositories"
	"github.com/google/uuid"
)

type SkillService struct {
	repo repositories.SkillRepository
}

func NewSkillService(repo repositories.SkillRepository) *SkillService {
	return &SkillService{repo: repo}
}

func (s *SkillService) CreateSkill(req models.CreateSkillRequest) (*models.Skill, error) {
	LastPracticed := (*time.Time)(nil)
	if req.LastPracticed != nil && *req.LastPracticed != "" {
		parsed, err := time.Parse("2006-01-02", *req.LastPracticed)
		if err != nil {
			return nil, errors.New("invalid date format")
		}
		LastPracticed = &parsed
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, errors.New("user_id inválido")
	}

	skill := &models.Skill{
		ID:         uuid.New(),
		UserID:     userID,
		Name:       req.Name,
		ProficiencyLevel:      req.ProficiencyLevel,
		LastPracticed: LastPracticed,
	}

	if err := s.repo.Create(skill); err != nil {
		return nil, err
	}										
	return skill, nil																
}

func (s *SkillService) GetSkillByID(id uuid.UUID) (*models.Skill, error) {
	return s.repo.GetByID(id)
}

func (s *SkillService) ListByUser(userID uuid.UUID) ([]*models.Skill, error) {
	return s.repo.ListByUser(userID)
}

func (s *SkillService) UpdateSkill(id uuid.UUID, req models.UpdateSkillRequest) (*models.Skill, error) {
    skill, err := s.repo.GetByID(id)
    if err != nil {
        return nil, err
    }
    if skill == nil {
		return nil, errors.New("skill not found")
    }

    if req.Name != nil {
        skill.Name = *req.Name
    }
    if req.ProficiencyLevel != nil {
        skill.ProficiencyLevel = *req.ProficiencyLevel
    }
    if req.LastPracticed != nil && *req.LastPracticed != "" {
        parsed, err := time.Parse("2006-01-02", *req.LastPracticed)
        if err != nil {
			return nil, errors.New("invalid date format")
        }
        skill.LastPracticed = &parsed
    }

    if err := s.repo.Update(skill); err != nil {
        return nil, err
    }
    return skill, nil
}

func (s *SkillService) DeleteSkill(id uuid.UUID) error {
	return s.repo.Delete(id)
}