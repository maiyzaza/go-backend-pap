package models

import (
	"time"

	models_Contract "PattayaAvenueProperty/models/Contract"
)

type RoomContractDocument struct {
	ID             uint       `gorm:"primaryKey"`
	DocumentID     uint       `gorm:"document_id"`
	RoomContractID uint       `gorm:"room_contråact_id"`
	IsActive       bool       `gorm:"is_active"`
	CreatedAt      time.Time  `gorm:"created_at"`
	UpdatedAt      *time.Time `gorm:"updated_at"`

	Document     Document                     `gorm:"foreignKey:DocumentID"`
	RoomContract models_Contract.RoomContract `gorm:"foreignKey:RoomContractID"`
}
