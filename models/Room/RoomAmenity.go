package models

import (
	"time"
)

type RoomAmenity struct {
	ID          uint       `json:"ID" gorm:"primaryKey"`
	RoomID      uint       `json:"RoomID" gorm:"room_id"`
	AmenityID   uint       `json:"AmenityID" gorm:"amenity_id"`
	Description string     `json:"Description" gorm:"description"`
	IsActive    bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt   time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt   *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Room    Room    `json:"Room" gorm:"foreignKey:RoomID"`
	Amenity Amenity `json:"Amenity" gorm:"foreignKey:AmenityID"`
}
