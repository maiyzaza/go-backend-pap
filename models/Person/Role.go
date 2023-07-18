package models

import (
	"time"
)

type Role struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"name"`
	Description *string   `gorm:"description"`
	IsActive    bool      `gorm:"is_active"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}
