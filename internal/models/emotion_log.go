package models

import (
	"github.com/google/uuid"
	"time"
)

type EmotionType string
type DayPeriod string

const (
    Happy    EmotionType = "happy"
    Sad      EmotionType = "sad"
    Neutral  EmotionType = "neutral"
    Calm     EmotionType = "calm"
    Anxious  EmotionType = "anxious"
    Stressed EmotionType = "stressed"
    Excited  EmotionType = "excited"
    Tired    EmotionType = "tired"
)

const (
    Morning   DayPeriod = "morning"
    Afternoon DayPeriod = "afternoon"
    Evening   DayPeriod = "evening"
)


type EmotionLog struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	Emotion   EmotionType `gorm:"type:emotion_type;not null" json:"emotion"`
	Period    DayPeriod   `gorm:"type:day_period;not null" json:"period"`
	LoggedAt  time.Time  `gorm:"type:timestamptz;default:current_timestamp" json:"logged_at"`
	Notes     *string    `json:"notes,omitempty"`
}

type CreateEmotionLogRequest struct {
	UserID   string  `json:"user_id" binding:"required,uuid"`
	Emotion  EmotionType  `json:"emotion" binding:"required"`
	Period   DayPeriod  `json:"period" binding:"required"`
	Notes    *string `json:"notes,omitempty"`
}

func (EmotionLog) TableName() string {
	return "emotion_log"
}