package models

import (
	models_Room "PattayaAvenueProperty/models/Room"
	"time"
)

type RoomDocument struct {
	ID          uint       `json:"ID" gorm:"primaryKey"`
	RoomID      uint       `json:"RoomID" gorm:"room_id"`
	DocumentUrl string     `json:"DocumentUrl" gorm:"document_url"`
	IsActive    bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt   time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt   *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Room models_Room.Room `json:"Room" gorm:"foreignKey:RoomID"`
}
