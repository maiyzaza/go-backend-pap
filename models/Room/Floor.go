package models

import (
	"time"
)

type Floor struct {
	ID          uint       `json:"ID" gorm:"primaryKey"`
	BuildingID  uint       `json:"BuildingID" gorm:"building_id"`
	FloorNumber string     `json:"FloorNumber" gorm:"floor_number"`
	IsActive    bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt   time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt   *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Building Building `json:"Building" gorm:"foreignKey:building_id"`
}
