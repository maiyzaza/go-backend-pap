package models

import (
	"time"
)

type RoomPicture struct {
	ID             int        `json:"ID" gorm:"primaryKey"`
	RoomID         int        `json:"RoomID" gorm:"room_id"`
	RoomPictureUrl string     `json:"RoomPictureUrl" gorm:"room_picture_url"`
	IsActive       bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt      time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt      *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Room Room `json:"Room" gorm:"foreignKey:RoomID"`
}
