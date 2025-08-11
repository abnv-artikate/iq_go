package models

import (
	"time"

	"gorm.io/gorm"
)

type TestResult struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	UserID         uint           `json:"user_id" gorm:"not null"`
	TestID         uint           `json:"test_id" gorm:"not null"`
	Score          int            `json:"score"`
	TotalQuestions int            `json:"total_questions"`
	TimeTaken      int            `json:"time_taken"` // in seconds
	StartedAt      time.Time      `json:"started_at"`
	CompletedAt    *time.Time     `json:"completed_at,omitempty"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`

	User    User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Test    Test     `json:"test,omitempty" gorm:"foreignKey:TestID"`
	Answers []Answer `json:"answers,omitempty" gorm:"foreignKey:TestResultID"`
}

type Answer struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	TestResultID uint           `json:"test_result_id" gorm:"not null"`
	QuestionID   uint           `json:"question_id" gorm:"not null"`
	UserAnswer   string         `json:"user_answer"`
	IsCorrect    bool           `json:"is_correct"`
	ResponseTime int            `json:"response_time"` // in milliseconds
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`

	TestResult TestResult `json:"test_result,omitempty" gorm:"foreignKey:TestResultID"`
	Question   Question   `json:"question,omitempty" gorm:"foreignKey:QuestionID"`
}
