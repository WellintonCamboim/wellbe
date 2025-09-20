package repositories

import (
	"errors"
	"github.com/WellintonCamboim/wellbe/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id uuid.UUID) (*models.Task, error)
	ListByUser(userID uuid.UUID) ([]*models.Task, error)
	Update(task *models.Task) error
	Delete(id uuid.UUID) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetByID(id uuid.UUID) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) ListByUser(userID uuid.UUID) ([]*models.Task, error) {
	var tasks []*models.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Task{}, "id = ?", id).Error
}

