package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PomodoroSession model
type PomodoroSession struct {
	SessionID    uuid.UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"session_id"`
	UserID       uuid.UUID     `gorm:"type:uuid;not null" json:"user_id"`
	CreatedAt    time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	WorkSessions []WorkSession `gorm:"foreignKey:SessionID" json:"work_sessions,omitempty"`
	BreakSessions []BreakSession `gorm:"foreignKey:SessionID" json:"break_sessions,omitempty"`
}

func (ps *PomodoroSession) BeforeCreate(tx *gorm.DB) (err error) {
	if ps.SessionID == uuid.Nil {
		ps.SessionID = uuid.New()
	}
	return
}