package models

import (
	"time"
)

type Building struct {
	ID           uint       `gorm:"primaryKey"`
	PlaceID      uint       `gorm:"place_id"`
	BuildingName string     `gorm:"building_name"`
	IsActive     bool       `gorm:"is_active"`
	CreatedAt    time.Time  `gorm:"created_at"`
	UpdatedAt    *time.Time `gorm:"updated_at"`
	Floors       []Floor    `gorm:"Floor;foreignKey:building_id"`

	Place Place `gorm:"foreignKey:place_id"`
}
