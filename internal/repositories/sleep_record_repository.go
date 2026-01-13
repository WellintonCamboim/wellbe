package repositories

import (
	"errors"

	"github.com/WellintonCamboim/wellbe/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SleepRecordRepository interface {
	Create(sleepRecord *models.SleepRecord) error
	GetByID(id uuid.UUID) (*models.SleepRecord, error)
	GetByUserAndDate(userID uuid.UUID, recordDate string) (*models.SleepRecord, error)
	ListByUser(userID uuid.UUID) ([]*models.SleepRecord, error)
	Update(sleepRecord *models.SleepRecord) error
	Delete(id uuid.UUID) error
}

type sleepRecordRepository struct {
	db *gorm.DB
}

func NewSleepRecordRepository(db *gorm.DB) SleepRecordRepository {
	return &sleepRecordRepository{db: db}
}

func (r *sleepRecordRepository) Create(sleepRecord *models.SleepRecord) error {
	return r.db.Create(sleepRecord).Error
}

func (r *sleepRecordRepository) GetByID(id uuid.UUID) (*models.SleepRecord, error) {
	var sleepRecord models.SleepRecord
	err := r.db.First(&sleepRecord, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &sleepRecord, nil
}

func (r *sleepRecordRepository) GetByUserAndDate(userID uuid.UUID, recordDate string) (*models.SleepRecord, error) {
	var sleepRecord models.SleepRecord
	err := r.db.Where("user_id = ? AND record_date = ?", userID, recordDate).First(&sleepRecord).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &sleepRecord, nil
}

func (r *sleepRecordRepository) ListByUser(userID uuid.UUID) ([]*models.SleepRecord, error) {
	var sleepRecords []*models.SleepRecord
	err := r.db.Where("user_id = ?", userID).Order("record_date DESC").Find(&sleepRecords).Error
	return sleepRecords, err
}

func (r *sleepRecordRepository) Update(sleepRecord *models.SleepRecord) error {
	return r.db.Save(sleepRecord).Error
}

func (r *sleepRecordRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.SleepRecord{}, "id = ?", id).Error
}
