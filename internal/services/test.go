package services

import (
	"strings"
	"time"

	"iq-go/internal/models"

	"gorm.io/gorm"
)

type TestService struct {
	db *gorm.DB
}

func NewTestService(db *gorm.DB) *TestService {
	return &TestService{db: db}
}

type SubmitAnswerRequest struct {
	QuestionID   uint   `json:"question_id" binding:"required"`
	UserAnswer   string `json:"user_answer"`
	ResponseTime int    `json:"response_time"`
}

func (s *TestService) GetQuestionsByTestID(testID uint) ([]models.Question, error) {
	var questions []models.Question
	err := s.db.Where("test_id = ?", testID).Order("order_index").Find(&questions).Error
	return questions, err
}

func (s *TestService) SubmitTest(userID, testID uint, answers []SubmitAnswerRequest, timeTaken int) (*models.TestResult, error) {
	// Create test result
	testResult := &models.TestResult{
		UserID:         userID,
		TestID:         testID,
		TimeTaken:      timeTaken,
		StartedAt:      time.Now().Add(-time.Duration(timeTaken) * time.Second),
		CompletedAt:    &[]time.Time{time.Now()}[0],
		TotalQuestions: len(answers),
	}

	err := s.db.Create(testResult).Error
	if err != nil {
		return nil, err
	}

	// Get all questions for this test
	var questions []models.Question
	err = s.db.Where("test_id = ?", testID).Find(&questions).Error
	if err != nil {
		return nil, err
	}

	// Create a map for quick question lookup
	questionMap := make(map[uint]*models.Question)
	for i := range questions {
		questionMap[questions[i].ID] = &questions[i]
	}

	score := 0
	var answerModels []models.Answer

	// Process each answer
	for _, answer := range answers {
		question, exists := questionMap[answer.QuestionID]
		if !exists {
			continue
		}

		isCorrect := s.evaluateAnswer(question, answer.UserAnswer)
		if isCorrect {
			score++
		}

		answerModel := models.Answer{
			TestResultID: testResult.ID,
			QuestionID:   answer.QuestionID,
			UserAnswer:   answer.UserAnswer,
			IsCorrect:    isCorrect,
			ResponseTime: answer.ResponseTime,
		}
		answerModels = append(answerModels, answerModel)
	}

	// Bulk create answers
	if len(answerModels) > 0 {
		err = s.db.Create(&answerModels).Error
		if err != nil {
			return nil, err
		}
	}

	// Update test result with score
	testResult.Score = score
	err = s.db.Save(testResult).Error
	if err != nil {
		return nil, err
	}

	// Load the complete result with relationships
	err = s.db.Preload("Answers").Preload("Test").First(testResult, testResult.ID).Error
	return testResult, err
}

func (s *TestService) evaluateAnswer(question *models.Question, userAnswer string) bool {
	correctAnswer := strings.TrimSpace(strings.ToLower(question.CorrectAnswer))
	userAnswer = strings.TrimSpace(strings.ToLower(userAnswer))

	switch question.QuestionType {
	case models.MultipleChoice:
		return correctAnswer == userAnswer
	case models.TextInput:
		// For text input, handle multiple correct answers separated by comma
		correctAnswers := strings.Split(correctAnswer, ",")
		for _, correct := range correctAnswers {
			if strings.TrimSpace(correct) == userAnswer {
				return true
			}
		}
		return false
	case models.NumberInput:
		return correctAnswer == userAnswer
	case models.KeySequence:
		return correctAnswer == userAnswer
	default:
		return correctAnswer == userAnswer
	}
}
