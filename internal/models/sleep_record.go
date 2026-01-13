package models

import (
	"time"

	"github.com/google/uuid"
)

type SleepRecord struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID         uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	RecordDate     time.Time `gorm:"type:date;not null;default:CURRENT_DATE" json:"record_date"`
	TargetHours    float64   `gorm:"type:decimal(3,1);not null;check:target_hours >= 4 AND target_hours <= 12" json:"target_hours"`
	ActualHours    float64   `gorm:"type:decimal(3,1);not null;check:actual_hours >= 0 AND actual_hours <= 24" json:"actual_hours"`
	QualityRating  *int      `gorm:"type:int;check:quality_rating >= 1 AND quality_rating <= 5" json:"quality_rating,omitempty"`
	WasInterrupted bool      `gorm:"type:bool;not null;default:false" json:"was_interrupted"`
	MedicationUsed bool      `gorm:"type:bool;not null;default:false" json:"medication_used"`
	Notes          *string   `gorm:"type:text" json:"notes,omitempty"`
	CreatedAt      time.Time `gorm:"type:timestamp;not null;default:now()" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

type CreateSleepRecordRequest struct {
	UserID         string  `json:"user_id" binding:"required,uuid"`
	RecordDate     string  `json:"record_date" binding:"required"`
	TargetHours    float64 `json:"target_hours" binding:"required,min=4,max=12"`
	ActualHours    float64 `json:"actual_hours" binding:"required,min=0,max=24"`
	QualityRating  *int    `json:"quality_rating,omitempty" binding:"omitempty,min=1,max=5"`
	WasInterrupted bool    `json:"was_interrupted"`
	MedicationUsed bool    `json:"medication_used"`
	Notes          *string `json:"notes,omitempty"`
}

type UpdateSleepRecordRequest struct {
	TargetHours    *float64 `json:"target_hours,omitempty" binding:"omitempty,min=4,max=12"`
	ActualHours    *float64 `json:"actual_hours,omitempty" binding:"omitempty,min=0,max=24"`
	QualityRating  *int     `json:"quality_rating,omitempty" binding:"omitempty,min=1,max=5"`
	WasInterrupted *bool    `json:"was_interrupted,omitempty"`
	MedicationUsed *bool    `json:"medication_used,omitempty"`
	Notes          *string  `json:"notes,omitempty"`
}

func (SleepRecord) TableName() string {
	return "sleep_record"
}
