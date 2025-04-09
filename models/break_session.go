package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompletionStatus string

const (
	StatusCompleted CompletionStatus = "completed"
	StatusSkipped   CompletionStatus = "skipped"
)

type CompletionMetric string

const (
	MetricDuration    CompletionMetric = "duration"
	MetricRepetitions CompletionMetric = "repetitions"
	MetricSimple      CompletionMetric = "simple"
)

type BreakSession struct {
	BreakID           uuid.UUID        `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"break_id"`
	SessionID         uuid.UUID        `gorm:"type:uuid;not null" json:"session_id"`
	TaskID            uuid.UUID        `gorm:"type:uuid" json:"task_id"`
	StartTime         time.Time        `gorm:"not null" json:"start_time"`
	EndTime           *time.Time       `json:"end_time"`
	IsShort           bool             `gorm:"default:false" json:"is_short"`
	Status            CompletionStatus `gorm:"type:text;not null" json:"status"`
	CompletionType    CompletionMetric `gorm:"type:text" json:"completion_type"`
	ActualDuration    *int             `json:"actual_duration"`
	ActualRepetitions *int             `json:"actual_repetitions"`
	Notes             string           `json:"notes"`
	CreatedAt         time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Task              *Task            `gorm:"foreignKey:TaskID" json:"task,omitempty"`
}

func (bs *BreakSession) BeforeCreate(tx *gorm.DB) (err error) {
	if bs.BreakID == uuid.Nil {
		bs.BreakID = uuid.New()
	}
	return
}
