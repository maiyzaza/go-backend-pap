package models

import (
	"time"
)

type RoomAmenity struct {
	ID          int        `gorm:"primaryKey"`
	RoomID      int        `gorm:"room_id"`
	AmenityID   int        `gorm:"amenity_id"`
	Description string     `gorm:"description"`
	IsActive    bool       `gorm:"is_active"`
	CreatedAt   time.Time  `gorm:"created_at"`
	UpdatedAt   *time.Time `gorm:"updated_at"`

	Room    Room    `gorm:"foreignKey:RoomID"`
	Amenity Amenity `gorm:"foreignKey:AmenityID"`
}
