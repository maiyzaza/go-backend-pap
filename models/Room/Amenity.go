package models

import (
	"time"
)

type Amenity struct {
	ID        uint       `json:"ID" gorm:"primaryKey"`
	Name      string     `json:"Name" gorm:"name"`
	IsActive  bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt *time.Time `json:"UpdatedAt" gorm:"updated_at"`
}
