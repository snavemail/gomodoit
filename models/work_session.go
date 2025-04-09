package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WorkSession struct {
	WorkID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"work_id"`
	SessionID  uuid.UUID `gorm:"type:uuid;not null" json:"session_id"`
	TagID      uuid.UUID `gorm:"type:uuid" json:"tag_id"`
	StartTime  time.Time `gorm:"not null" json:"start_time"`
	Duration   int       `gorm:"not null" json:"duration"` // In seconds
	Notes      string    `json:"notes"`
	Tag        *Tag      `gorm:"foreignKey:TagID" json:"tag,omitempty"`
}

func (ws *WorkSession) BeforeCreate(tx *gorm.DB) (err error) {
	if ws.WorkID == uuid.Nil {
		ws.WorkID = uuid.New()
	}
	return
}


