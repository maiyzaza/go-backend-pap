package models

import (
	"time"

	models_Room "PattayaAvenueProperty/models/Room"
)

type SecondaryContract struct {
	ID                     uint       `gorm:"primaryKey"`
	RoomID                 uint       `gorm:"room_id"`
	RoomContractID         uint       `gorm:"room_contract_id"`
	CheckInDate            *time.Time `gorm:"check_in_date"`
	CheckOutDate           *time.Time `gorm:"check_out_date"`
	CheckInElectricNumber  *float32   `gorm:"electric_number"`
	CheckInWaterNumber     *float32   `gorm:"water_number"`
	CheckOutElectricNumber *float32   `gorm:"electric_number"`
	CheckOutWaterNumber    *float32   `gorm:"water_number"`
	IsClosed               bool       `gorm:"is_closed"`
	IsActive               bool       `gorm:"is_active"`
	CreatedAt              time.Time  `gorm:"created_at"`
	UpdatedAt              *time.Time `gorm:"updated_at"`

	Room         models_Room.Room `gorm:"foreignKey:RoomID"`
	RoomContract RoomContract     `gorm:"foreignKey:RoomContractID"`
}
