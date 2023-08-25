package models

import (
	"time"
)

type RoomPicture struct {
	ID             uint       `json:"ID" gorm:"primaryKey"`
	RoomID         uint       `json:"RoomID" gorm:"room_id"`
	RoomPictureUrl string     `json:"RoomPictureUrl" gorm:"room_picture_url"`
	IsActive       bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt      time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt      *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Room Room `json:"Room" gorm:"foreignKey:RoomID"`
}
