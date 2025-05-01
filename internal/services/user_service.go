package services

import (
	"errors"
	"time"

	"github.com/WellintonCamboim/wellbe/internal/models"
	"github.com/WellintonCamboim/wellbe/internal/repositories"
)

type UserService struct {
    repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req models.CreateUserRequest) (*models.User, error) {
    exists, err := s.repo.EmailExists(req.Email)  // Corrigido: s.repo
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, errors.New("email já cadastrado")
    }

    birthDate, err := time.Parse("2006-01-02", req.BirthDate)
    if err != nil {
        return nil, errors.New("formato de data inválido")
    }

    user := &models.User{
        Email:          req.Email,
        BirthDate:      birthDate,
        Phone:          req.Phone,
        Profession:     req.Profession,
        EducationLevel: req.EducationLevel,
    }

    if err := s.repo.Create(user); err != nil {  // Corrigido: s.repo
        return nil, err
    }

    return user, nil
}

// Add to UserService
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
    user, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }
    if user == nil {
        return nil, errors.New("user not found")
    }
    return user, nil
}