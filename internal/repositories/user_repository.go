package repositories

import (
	"errors"

	"github.com/WellintonCamboim/wellbe/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
    Create(user *models.User) error
    EmailExists(email string) (bool, error)
    FindByID(id uint) (*models.User, error)
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

// Implement the FindByID method in userRepository
func (r *userRepository) FindByID(id uint) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil // User not found
        }
        return nil, err
    }
    return &user, nil
}