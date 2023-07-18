package models

import (
	"time"
)

type Floor struct {
	ID          uint       `gorm:"primaryKey"`
	BuildingID  uint       `gorm:"building_id"`
	FloorNumber string     `gorm:"floor_number"`
	IsActive    bool       `gorm:"is_active"`
	CreatedAt   time.Time  `gorm:"created_at"`
	UpdatedAt   *time.Time `gorm:"updated_at"`

	Building Building `gorm:"foreignKey:building_id"`
}
