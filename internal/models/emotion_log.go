package models

import (
	"github.com/google/uuid"
	"time"
)

type EmotionLog struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	Emotion   string     `gorm:"type:emotion_type;not null" json:"emotion"`
	Period    string     `gorm:"type:day_period;not null" json:"period"`
	LoggedAt  time.Time  `gorm:"type:timestamptz;default:current_timestamp" json:"logged_at"`
	Notes     *string    `json:"notes,omitempty"`
}

type CreateEmotionLogRequest struct {
	UserID   string  `json:"user_id" binding:"required,uuid"`
	Emotion  string  `json:"emotion" binding:"required"`
	Period   string  `json:"period" binding:"required"`
	Notes    *string `json:"notes,omitempty"`
}

func (EmotionLog) TableName() string {
	return "emotion_log"
}