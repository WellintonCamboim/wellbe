package services

import (
	"errors"
	"time"
	"github.com/WellintonCamboim/wellbe/internal/models"
	"github.com/WellintonCamboim/wellbe/internal/repositories"
	"github.com/google/uuid"
)

type EmotionLogService struct {
	repo repositories.EmotionLogRepository
}

func NewEmotionLogService(repo repositories.EmotionLogRepository) *EmotionLogService {
	return &EmotionLogService{repo: repo}
}

func (s *EmotionLogService) CreateEmotionLog(req *models.CreateEmotionLogRequest) (*models.EmotionLog, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return nil, errors.New("user_id inválido")
	}

	emotionLog := &models.EmotionLog{
		ID:       uuid.New(),
		UserID:   userID,
		Emotion:  req.Emotion,
		Period:   req.Period,
		LoggedAt: time.Now(),
		Notes:    req.Notes,
	}

	if err := s.repo.Create(emotionLog); err != nil {
		return nil, err
	}

	return emotionLog, nil
}


