package model

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Schema struct {
	ID          uint            `gorm:"primaryKey;autoIncrement"`
	Level       string          `gorm:"type:varchar(50);not null"`
	ServiceName *string         `gorm:"type:varchar(50)"` // Nullable for global levels
	SchemaData  json.RawMessage `gorm:"type:jsonb;not null"`
	CreatedAt   time.Time       `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time       `gorm:"default:current_timestamp"`
	DeletedAt   gorm.DeletedAt  `gorm:"index"` // Soft delete support
}
