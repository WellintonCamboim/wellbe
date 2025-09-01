package models

import (
	"time"
	"github.com/google/uuid"
)

type Task struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID  `gorm:"type:uuid;not null" json:"user_id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	DueDate      *time.Time `gorm:"column:due_date" json:"due_date,omitempty"`
	IsCompleted  bool       `gorm:"column:is_completed" json:"is_completed"`
	CompletedAt  *time.Time `gorm:"column:completed_at" json:"completed_at,omitempty"`
	CreatedAt    time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

type CreateTaskRequest struct {
	UserID      string  `json:"user_id" binding:"required,uuid"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required"`
	DueDate     *string `json:"due_date,omitempty" binding:"omitempty,datetime=2006-01-02"`
}

func (Task) TableName() string {
	return "task"
}