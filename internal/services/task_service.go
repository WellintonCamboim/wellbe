package services

import (
    "errors"
    "time"

    "github.com/WellintonCamboim/wellbe/internal/models"
    "github.com/WellintonCamboim/wellbe/internal/repositories"
    "github.com/google/uuid"
)

type TaskService struct {
    repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) *TaskService {
    return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(req models.CreateTaskRequest) (*models.Task, error) {
    dueDate := (*time.Time)(nil)
    if req.DueDate != nil && *req.DueDate != "" {
        parsed, err := time.Parse("2006-01-02", *req.DueDate)
        if err != nil {
            return nil, errors.New("formato de data inválido")
        }
        dueDate = &parsed
    }

    userID, err := uuid.Parse(req.UserID)
    if err != nil {
        return nil, errors.New("user_id inválido")
    }

    task := &models.Task{
        ID:          uuid.New(),
        UserID:      userID,
        Title:       req.Title,
        Description: req.Description,
        DueDate:     dueDate,
        IsCompleted: false,
    }

    if err := s.repo.Create(task); err != nil {
        return nil, err
    }

    return task, nil
}

func (s *TaskService) GetTaskByID(id uuid.UUID) (*models.Task, error) {
    task, err := s.repo.GetByID(id)
    if err != nil {
        return nil, err
    }
    if task == nil {
        return nil, errors.New("task not found")
    }
    return task, nil
}

func (s *TaskService) ListTasksByUser(userID uuid.UUID) ([]*models.Task, error) {
    return s.repo.ListByUser(userID)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
    return s.repo.Update(task)
}

func (s *TaskService) DeleteTask(id uuid.UUID) error {
    return s.repo.Delete(id)
}