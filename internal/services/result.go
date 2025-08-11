package services

import (
	"iq-go/internal/models"

	"gorm.io/gorm"
)

type ResultService struct {
	db *gorm.DB
}

func NewResultService(db *gorm.DB) *ResultService {
	return &ResultService{db: db}
}

func (s *ResultService) GetResultsByUserID(userID uint) ([]models.TestResult, error) {
	var results []models.TestResult
	err := s.db.Where("user_id = ?", userID).
		Preload("Test").
		Order("created_at DESC").
		Find(&results).Error
	return results, err
}

func (s *ResultService) GetResultByID(resultID, userID uint) (*models.TestResult, error) {
	var result models.TestResult
	err := s.db.Where("id = ? AND user_id = ?", resultID, userID).
		Preload("Test").
		Preload("Answers").
		Preload("Answers.Question").
		First(&result).Error
	return &result, err
}

func (s *ResultService) CreateResult(result *models.TestResult) error {
	return s.db.Create(result).Error
}

func (s *ResultService) UpdateResult(result *models.TestResult) error {
	return s.db.Save(result).Error
}
