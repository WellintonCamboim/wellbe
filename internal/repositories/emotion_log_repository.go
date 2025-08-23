package repositories

import (
	"github.com/WellintonCamboim/wellbe/internal/models"
	"gorm.io/gorm"
)

type EmotionLogRepository interface {
	Create(emotion_log *models.EmotionLog) error
}

type emotionLogRepository struct {
	db *gorm.DB
}

func NewEmotionLogRepository(db *gorm.DB) EmotionLogRepository {
	return &emotionLogRepository{db: db}
}

func (r *emotionLogRepository) Create(emotion_log *models.EmotionLog) error {
	return r.db.Create(emotion_log).Error
}