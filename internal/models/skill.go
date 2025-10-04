package models

import (
	"time"
	"github.com/google/uuid"
)

type Skill struct {
	ID              uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID          uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	Name            string     `gorm:"type:varchar(100);not null" json:"name"`
	ProficiencyLevel int       `gorm:"type:int;not null" json:"proficiency_level"`
	Category        *string    `gorm:"type:varchar(50)" json:"category,omitempty"`
	LastPracticed   *time.Time `gorm:"type:date" json:"last_practiced,omitempty"`
	CreatedAt       time.Time  `gorm:"type:timestamp;not null;default:now()" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type CreateSkillRequest struct {
	UserID          string `json:"user_id" binding:"required,uuid"`
	Name            string `json:"name" binding:"required"`
	ProficiencyLevel int    `json:"proficiency_level" binding:"required,min=1,max=10"`
	Category         *string `json:"category,omitempty" binding:"omitempty,max=50"`
    LastPracticed    *string `json:"last_practiced,omitempty" binding:"omitempty,datetime=2006-01-02"`
}

type UpdateSkillRequest struct {
    Name             *string `json:"name,omitempty" binding:"omitempty,max=100"`
    ProficiencyLevel *int    `json:"proficiency_level,omitempty" binding:"omitempty,min=1,max=10"`
    Category         *string `json:"category,omitempty" binding:"omitempty,max=50"`
    LastPracticed    *string `json:"last_practiced,omitempty" binding:"omitempty,datetime=2006-01-02"`
}

func (Skill) TableName() string {
	return "skill"
}