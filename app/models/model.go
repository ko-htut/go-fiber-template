package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time      `json:"updated_at"`
	UpdatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
