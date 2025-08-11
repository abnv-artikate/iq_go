package models

import (
	"time"

	"gorm.io/gorm"
)

type QuestionType string

const (
	MultipleChoice QuestionType = "multiple_choice"
	TextInput      QuestionType = "text_input"
	NumberInput    QuestionType = "number_input"
	KeySequence    QuestionType = "key_sequence"
)

type Category string

const (
	AnalyticalReasoning Category = "analytical_reasoning"
	WorkingMemory       Category = "working_memory"
	ProcessingSpeed     Category = "processing_speed"
	AttentionFocus      Category = "attention_focus"
	EmotionalRegulation Category = "emotional_regulation"
)

type Question struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	TestID        uint           `json:"test_id"`
	QuestionText  string         `json:"question_text" gorm:"type:text;not null"`
	QuestionType  QuestionType   `json:"question_type" gorm:"not null"`
	Category      Category       `json:"category" gorm:"not null"`
	Options       string         `json:"options,omitempty" gorm:"type:text"` // JSON array for multiple choice
	CorrectAnswer string         `json:"correct_answer" gorm:"not null"`
	TimeLimit     int            `json:"time_limit"`   // in seconds
	DisplayTime   int            `json:"display_time"` // in seconds for memory questions
	OrderIndex    int            `json:"order_index"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`

	Test Test `json:"test,omitempty" gorm:"foreignKey:TestID"`
}
