package models

import (
	"time"

	"gorm.io/gorm"
)

type Test struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Duration    int            `json:"duration"` // in minutes
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`

	Questions []Question `json:"questions,omitempty" gorm:"foreignKey:TestID"`
}
