package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"user_id"`
	DeviceID      string    `gorm:"unique" json:"device_id"`
	Email         string    `gorm:"unique" json:"email"`
	EmailVerified bool      `gorm:"default:false" json:"email_verified"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UserID == uuid.Nil {
		u.UserID = uuid.New()
	}
	return
}