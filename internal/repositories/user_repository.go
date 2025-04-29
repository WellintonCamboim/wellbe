package repositories

import (
    "github.com/WellintonCamboim/wellbe/internal/models"
    "gorm.io/gorm"
)

type UserRepository interface {
    Create(user *models.User) error
    EmailExists(email string) (bool, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) EmailExists(email string) (bool, error) {
    var count int64
    err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
    return count > 0, err
}