package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskType string

const (
	TaskTypeTimed     TaskType = "timed"
	TaskTypeCountable TaskType = "countable"
	TaskTypeSimple    TaskType = "simple"
)

type Task struct {
	TaskID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"task_id"`
	FolderID         uuid.UUID `gorm:"type:uuid" json:"folder_id"`
	UserID           uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Title            string    `gorm:"size:100;not null" json:"title"`
	Description      string    `json:"description"`
	IsShort          bool      `gorm:"default:false" json:"is_short"`
	TaskType         TaskType  `gorm:"type:text;not null;default:'simple'" json:"task_type"`
	TargetDuration   int       `json:"target_duration"`
	TargetRepetitions int      `json:"target_repetitions"`
	IsActive         bool      `gorm:"default:true" json:"is_active"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	if t.TaskID == uuid.Nil {
		t.TaskID = uuid.New()
	}
	return
}
