package models

import (
	"time"
)

type Place struct {
	ID        uint       `gorm:"primaryKey"`
	PlaceName string     `gorm:"place_name"`
	IsActive  bool       `gorm:"is_active"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt *time.Time `gorm:"updated_at"`
	Buildings []Building `gorm:"building;foreignKey:place_id"`
}
