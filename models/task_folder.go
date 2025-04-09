package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TaskFolder model
type TaskFolder struct {
	FolderID  uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"folder_id"`
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Tasks     []Task    `gorm:"foreignKey:FolderID" json:"tasks,omitempty"`
}

func (tf *TaskFolder) BeforeCreate(tx *gorm.DB) (err error) {
	if tf.FolderID == uuid.Nil {
		tf.FolderID = uuid.New()
	}
	return
}