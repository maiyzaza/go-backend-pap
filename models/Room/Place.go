package models

import (
	"time"
)

type Place struct {
	ID        uint       `json:"ID" gorm:"primaryKey"`
	PlaceName string     `json:"PlaceName" gorm:"place_name"`
	IsActive  bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt *time.Time `json:"UpdatedAt" gorm:"updated_at"`
	Buildings []Building `json:"Buildings" gorm:"building;foreignKey:place_id"`
}
