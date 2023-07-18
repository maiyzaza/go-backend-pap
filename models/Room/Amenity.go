package models

import (
	"time"
)

type Amenity struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"name"`
	IsActive  bool       `gorm:"is_active"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"`
}
