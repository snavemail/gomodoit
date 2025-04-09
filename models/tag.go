package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type Tag struct {
	TagID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"tag_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Name      string    `gorm:"size:50;not null" json:"name"`
	Color     string    `gorm:"size:7;default:'#FF6B6B'" json:"color"` // Hex color code
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	if t.TagID == uuid.Nil {
		t.TagID = uuid.New()
	}
	return
}