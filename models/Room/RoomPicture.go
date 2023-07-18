package models

import (
	"time"
)

type RoomPicture struct {
	ID             int        `gorm:"primaryKey"`
	RoomID         int        `gorm:"room_id"`
	RoomPictureUrl string     `gorm:"room_picture_url"`
	IsActive       bool       `gorm:"is_active"`
	CreatedAt      time.Time  `gorm:"created_at"`
	UpdatedAt      *time.Time `gorm:"updated_at"`

	Room Room `gorm:"foreignKey:RoomID"`
}
