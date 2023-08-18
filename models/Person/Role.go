package models

import (
	"time"
)

type Role struct {
	ID          uint      `json:"ID" gorm:"primaryKey"`
	Name        string    `json:"Name" gorm:"name"`
	Description *string   `json:"Description" gorm:"description"`
	IsActive    bool      `json:"IsActive" gorm:"is_active"`
	CreatedAt   time.Time `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt   time.Time `json:"UpdatedAt" gorm:"updated_at"`
}
