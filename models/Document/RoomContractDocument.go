package models

import (
	"time"

	models_Contract "PattayaAvenueProperty/models/Contract"
)

type RoomContractDocument struct {
	ID             uint       `json:"ID" gorm:"primaryKey"`
	DocumentID     uint       `json:"DocumentID" gorm:"document_id"`
	RoomContractID uint       `json:"RoomContractID" gorm:"room_contract_id"`
	IsActive       bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt      time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt      *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Document     Document                     `json:"Document" gorm:"foreignKey:DocumentID"`
	RoomContract models_Contract.RoomContract `json:"RoomContract" gorm:"foreignKey:RoomContractID"`
}
