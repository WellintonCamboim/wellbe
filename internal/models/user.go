package models

import "time"

// User represents a user in the system
type User struct {
    ID             uint      `json:"id" gorm:"primaryKey" example:"1"`
    Email          string    `json:"email" gorm:"unique;not null" example:"user@example.com"`
    BirthDate      time.Time `json:"birth_date" gorm:"not null" example:"1990-01-01T00:00:00Z"`
    Phone          *string   `json:"phone,omitempty" gorm:"size:20" example:"+5511999999999"`
    Profession     *string   `json:"profession,omitempty" gorm:"size:100" example:"Software Engineer"`
    EducationLevel *string   `json:"education_level,omitempty" gorm:"size:100" example:"Bachelor's Degree"`
    CreatedAt      time.Time `json:"created_at" gorm:"default:now()" example:"2023-01-01T00:00:00Z"`
    UpdatedAt      time.Time `json:"updated_at" gorm:"default:now()" example:"2023-01-01T00:00:00Z"`
}

// CreateUserRequest defines the structure for user creation
type CreateUserRequest struct {
    Email          string  `json:"email" validate:"required,email" example:"user@example.com"`
    BirthDate      string  `json:"birth_date" validate:"required,datetime=2006-01-02" example:"1990-01-01"`
    Phone          *string `json:"phone,omitempty" validate:"omitempty,min=10" example:"+5511999999999"`
    Profession     *string `json:"profession,omitempty" example:"Software Engineer"`
    EducationLevel *string `json:"education_level,omitempty" example:"Bachelor's Degree"`
}

// TableName defines the table name in the database
func (User) TableName() string {
    return "user"
}