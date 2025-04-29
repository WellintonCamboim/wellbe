package models

import "time"

type User struct {
    ID             uint      `json:"id" gorm:"primaryKey"`
    Email          string    `json:"email" gorm:"unique;not null"`
    BirthDate      time.Time `json:"birth_date" gorm:"not null"`
    Phone          *string   `json:"phone,omitempty" gorm:"size:20"`
    Profession     *string   `json:"profession,omitempty" gorm:"size:100"`
    EducationLevel *string   `json:"education_level,omitempty" gorm:"size:100"`
    CreatedAt      time.Time `json:"created_at" gorm:"default:now()"`
    UpdatedAt      time.Time `json:"updated_at" gorm:"default:now()"`
}

type CreateUserRequest struct {
    Email          string  `json:"email" validate:"required,email"`
    BirthDate      string  `json:"birth_date" validate:"required,datetime=2006-01-02"`
    Phone          *string `json:"phone,omitempty" validate:"omitempty,min=10"`
    Profession     *string `json:"profession,omitempty"`
    EducationLevel *string `json:"education_level,omitempty"`
}

// Define o nome da tabela como "user" (singular)
func (User) TableName() string {
    return "user"
}