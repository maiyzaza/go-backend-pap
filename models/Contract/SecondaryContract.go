package models

import (
	"time"

	models_Room "PattayaAvenueProperty/models/Room"
)

type SecondaryContract struct {
	ID                     uint       `json:"ID" gorm:"primaryKey"`
	RoomID                 uint       `json:"RoomID" gorm:"room_id"`
	RoomContractID         uint       `json:"RoomContractID" gorm:"room_contract_id"`
	CheckInDate            *time.Time `json:"CheckInDate" gorm:"check_in_date"`
	CheckOutDate           *time.Time `json:"CheckOutDate" gorm:"check_out_date"`
	CheckInElectricNumber  *float32   `json:"CheckInElectricNumber" gorm:"electric_number"`
	CheckInWaterNumber     *float32   `json:"CheckInWaterNumber" gorm:"water_number"`
	CheckOutElectricNumber *float32   `json:"CheckOutElectricNumber" gorm:"electric_number"`
	CheckOutWaterNumber    *float32   `json:"CheckOutWaterNumber" gorm:"water_number"`
	IsClosed               bool       `json:"IsClosed" gorm:"is_closed"`
	IsActive               bool       `json:"IsActive" gorm:"is_active"`
	CreatedAt              time.Time  `json:"CreatedAt" gorm:"created_at"`
	UpdatedAt              *time.Time `json:"UpdatedAt" gorm:"updated_at"`

	Room         models_Room.Room `json:"Room" gorm:"foreignKey:RoomID"`
	RoomContract RoomContract     `json:"RoomContract" gorm:"foreignKey:RoomContractID"`
}
