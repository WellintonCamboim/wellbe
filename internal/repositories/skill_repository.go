package repositories

import (
	"errors"
	"github.com/WellintonCamboim/wellbe/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SkillRepository interface {
	Create(skill *models.Skill) error
	GetByID(id uuid.UUID) (*models.Skill, error)
	ListByUser(userID uuid.UUID) ([]*models.Skill, error)
	Update(skill *models.Skill) error
	Delete(id uuid.UUID) error
}

type skillRepository struct {
	db *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillRepository{db: db}
}

func (r *skillRepository) Create(skill *models.Skill) error {
	return r.db.Create(skill).Error
}

func (r *skillRepository) GetByID(id uuid.UUID) (*models.Skill, error) {
	var skill models.Skill
	err := r.db.First(&skill, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &skill, nil
}

func (r *skillRepository) ListByUser(userID uuid.UUID) ([]*models.Skill, error) {
	var skills []*models.Skill
	err := r.db.Where("user_id = ?", userID).Find(&skills).Error
	return skills, err
}

func (r *skillRepository) Update(skill *models.Skill) error {
	return r.db.Save(skill).Error
}

func (r *skillRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Skill{}, "id = ?", id).Error
}		