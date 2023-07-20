package models

import (
	"time"

	models_Room "PattayaAvenueProperty/models/Room"
)

type RoomContract struct {
	ID                     uint       `gorm:"primaryKey"`
	RoomID                 uint       `gorm:"room_id"`
	ContractName           string     `gorm:"contract_name"`
	StartContractDate      time.Time  `gorm:"start_contract_date"`
	EndContractDate        time.Time  `gorm:"end_contract_date"`
	Rental                 float32    `gorm:"rental"`
	Deposit                float32    `gorm:"deposit"`
	CheckInElectricNumber  *int       `gorm:"check_in_electric_number"`
	CheckInWaterNumber     *int       `gorm:"check_in_water_number"`
	CheckOutElectricNumber *int       `gorm:"check_out_electric_number"`
	CheckOutWaterNumber    *int       `gorm:"check_out_water_number"`
	CheckInDate            *time.Time `gorm:"check_in_date"`
	CheckOutDate           *time.Time `gorm:"check_out_date"`
	IsClosed               bool       `gorm:"is_closed"`
	IsActive               bool       `gorm:"is_active"`
	CreatedAt              time.Time  `gorm:"created_at"`
	UpdatedAt              *time.Time `gorm:"updated_at"`

	Room models_Room.Room `gorm:"foreignKey:RoomID"`
}
