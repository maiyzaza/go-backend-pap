package models

import (
	"time"
)

type Building struct {
	ID           uint       `json:"ID" gorm:"primaryKey"`
	PlaceID      uint       `json:"PlaceID" gorm:"place_id"`
	BuildingName string     `json:"BuildingName" gorm:"building_name"`
	IsActive     bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt    time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt    *time.Time `json:"UpdatedAt" gorm:"updated_at"`
	Floors       []Floor    `json:"Floors" gorm:"Floor;foreignKey:building_id"`

	Place Place `json:"Place" gorm:"foreignKey:place_id"`
}
