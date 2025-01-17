package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Configuration struct {
	ID         uint            `gorm:"primaryKey;autoIncrement"`
	LevelName  string          `gorm:"type:varchar(50);not null"`
	EntityID   string          `gorm:"type:varchar(255);not null"`
	ConfigData json.RawMessage `gorm:"type:jsonb;not null"` // JSONB type requires GORM plugin
	CreatedAt  time.Time       `gorm:"default:current_timestamp"`
	UpdatedAt  time.Time       `gorm:"default:current_timestamp"`
	DeletedAt  gorm.DeletedAt  `gorm:"index"` // Soft delete support
}
